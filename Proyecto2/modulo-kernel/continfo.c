#include <linux/module.h>
#include <linux/kernel.h>
#include <linux/init.h>
#include <linux/proc_fs.h>
#include <linux/seq_file.h>
#include <linux/sched/signal.h>
#include <linux/sched/mm.h>
#include <linux/mm.h>
#include <linux/jiffies.h>
#include <linux/sched.h>
#include <linux/slab.h>

#define PROC_NAME "continfo_so1_201700399"
#define MAX_CMDLINE 256

MODULE_LICENSE("GPL");
MODULE_AUTHOR("201700399");
MODULE_DESCRIPTION("Continfo: procesos generales y de contenedores con métricas");

static char *my_get_cmdline(struct task_struct *task)
{
    struct mm_struct *mm;
    char *buf;
    unsigned long arg_start, arg_end;
    int len;
    int i; /* C90: declarar antes del for */

    buf = kmalloc(MAX_CMDLINE, GFP_KERNEL);
    if (!buf)
        return NULL;

    mm = get_task_mm(task);
    if (!mm) {
        kfree(buf);
        return NULL;
    }

    down_read(&mm->mmap_sem);
    arg_start = mm->arg_start;
    arg_end   = mm->arg_end;
    up_read(&mm->mmap_sem);

    len = arg_end > arg_start ? (int)(arg_end - arg_start) : 0;
    if (len <= 0 || len >= MAX_CMDLINE)
        len = MAX_CMDLINE - 1;

    if (access_process_vm(task, arg_start, buf, len, 0) != len) {
        mmput(mm);
        kfree(buf);
        return NULL;
    }

    buf[len] = '\0';

    for (i = 0; i < len; i++)
        if (buf[i] == '\0')
            buf[i] = ' ';

    mmput(mm);
    return buf;
}

static int is_container_like(const char *comm)
{
    if (!comm) return 0;
    if (strnstr(comm, "container", strlen(comm))) return 1;
    if (strnstr(comm, "containerd", strlen(comm))) return 1;
    if (strnstr(comm, "docker", strlen(comm))) return 1;
    if (strnstr(comm, "runc", strlen(comm))) return 1;
    if (strnstr(comm, "shim", strlen(comm))) return 1;
    if (strnstr(comm, "dockerd", strlen(comm))) return 1;
    if (strnstr(comm, "docker-proxy", strlen(comm))) return 1;
    return 0;
}

static int continfo_show(struct seq_file *m, void *v)
{
    struct task_struct *task;
    struct sysinfo si;
    unsigned long total_kb;
    unsigned long jiffies_now;
    unsigned int nr_cpus;
    int first_proc;
    struct mm_struct *mm;
    unsigned long vsz_kb;
    unsigned long rss_kb;
    unsigned long mem_percent_x100;
    unsigned long cpu_percent_x100;
    unsigned long long total_time;
    char *cmdline;

    int i; /* para C90 */

    si_meminfo(&si);
    total_kb = (si.totalram * PAGE_SIZE) >> 10;
    jiffies_now = get_jiffies_64();
    nr_cpus = num_online_cpus();
    first_proc = 1;

    seq_puts(m, "{\n  \"procesos\": [\n");

    for_each_process(task) {

        vsz_kb = 0;
        rss_kb = 0;
        mem_percent_x100 = 0;
        cpu_percent_x100 = 0;
        total_time = 0;
        cmdline = NULL;

        mm = task->mm;
        if (mm) {
            vsz_kb = (mm->total_vm << (PAGE_SHIFT - 10));
            rss_kb = (get_mm_rss(mm) << (PAGE_SHIFT - 10));
        }

        if (total_kb > 0)
            mem_percent_x100 = (rss_kb * 10000) / total_kb;

        total_time = (unsigned long long)task->utime + (unsigned long long)task->stime;
        if (jiffies_now > 0) {
            cpu_percent_x100 = (total_time * 10000ULL) / jiffies_now;
            if (nr_cpus > 0)
                cpu_percent_x100 /= nr_cpus;
        }

        if (is_container_like(task->comm))
            cmdline = my_get_cmdline(task);

        if (!first_proc)
            seq_puts(m, ",\n");
        first_proc = 0;

        if (cmdline) {
            seq_printf(m,
                "    {\"pid\": %d, \"name\": \"%s\", \"cmdline\": \"%s\", \"vsz_kb\": %lu, \"rss_kb\": %lu, \"mem_percent\": %lu.%02lu, \"cpu_percent\": %lu.%02lu}",
                task->pid,
                task->comm,
                cmdline,
                vsz_kb,
                rss_kb,
                mem_percent_x100 / 100, mem_percent_x100 % 100,
                cpu_percent_x100 / 100, cpu_percent_x100 % 100
            );
            kfree(cmdline);
        } else {
            seq_printf(m,
                "    {\"pid\": %d, \"name\": \"%s\", \"vsz_kb\": %lu, \"rss_kb\": %lu, \"mem_percent\": %lu.%02lu, \"cpu_percent\": %lu.%02lu, \"state\": %ld}",
                task->pid,
                task->comm,
                vsz_kb,
                rss_kb,
                mem_percent_x100 / 100, mem_percent_x100 % 100,
                cpu_percent_x100 / 100, cpu_percent_x100 % 100,
                task->state
            );
        }
    }

    seq_puts(m, "\n  ]\n}\n");
    return 0;
}

static int continfo_open(struct inode *inode, struct file *file)
{
    return single_open(file, continfo_show, NULL);
}

static const struct file_operations continfo_fops = {
    .owner = THIS_MODULE,
    .open = continfo_open,
    .read = seq_read,
    .llseek = seq_lseek,
    .release = single_release,
};

static int __init continfo_init(void)
{
    if (!proc_create(PROC_NAME, 0444, NULL, &continfo_fops)) {
        pr_err("continfo: fallo creando /proc/%s\n", PROC_NAME);
        return -ENOMEM;
    }
    pr_info("continfo: modulo cargado\n");
    return 0;
}

static void __exit continfo_exit(void)
{
    remove_proc_entry(PROC_NAME, NULL);
    pr_info("continfo: modulo descargado\n");
}

module_init(continfo_init);
module_exit(continfo_exit);
