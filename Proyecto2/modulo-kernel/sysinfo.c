#include <linux/module.h>
#include <linux/kernel.h>
#include <linux/init.h>
#include <linux/proc_fs.h>
#include <linux/seq_file.h>
#include <linux/mm.h>
#include <linux/sched/signal.h>

#define PROC_NAME "sysinfo_so1_201700399"

static int sysinfo_show(struct seq_file *m, void *v)
{
    struct sysinfo si;
    struct task_struct *task;
    unsigned long total_kb;
    unsigned long free_kb;
    unsigned long used_kb;
    int first;

    si_meminfo(&si);

    total_kb = (si.totalram * PAGE_SIZE) >> 10;
    free_kb  = (si.freeram * PAGE_SIZE) >> 10;
    used_kb  = total_kb - free_kb;

    seq_puts(m, "{\n");
    seq_printf(m, "  \"memory\": {\n");
    seq_printf(m, "    \"total_kb\": %lu,\n", total_kb);
    seq_printf(m, "    \"free_kb\": %lu,\n", free_kb);
    seq_printf(m, "    \"used_kb\": %lu\n", used_kb);
    seq_puts(m, "  },\n");

    seq_puts(m, "  \"processes\": [\n");
    first = 1;
    for_each_process(task) {
        unsigned long vsz = 0;
        unsigned long rss = 0;

        if (task->mm) {
            vsz = (task->mm->total_vm << (PAGE_SHIFT - 10));
            rss = (get_mm_rss(task->mm) << (PAGE_SHIFT - 10));
        }

        if (!first)
            seq_puts(m, ",\n");
        first = 0;

        seq_printf(m,
            "    {\"pid\": %d, \"name\": \"%s\", \"vsz_kb\": %lu, \"rss_kb\": %lu, \"state\": %ld}",
            task->pid,
            task->comm,
            vsz,
            rss,
            task->state
        );
    }

    seq_puts(m, "\n  ]\n}\n");
    return 0;
}

static int sysinfo_open(struct inode *inode, struct file *file)
{
    return single_open(file, sysinfo_show, NULL);
}

static const struct file_operations sysinfo_ops = {
    .owner   = THIS_MODULE,
    .open    = sysinfo_open,
    .read    = seq_read,
    .llseek  = seq_lseek,
    .release = single_release,
};

static int __init sysinfo_init(void)
{
    proc_create(PROC_NAME, 0, NULL, &sysinfo_ops);
    pr_info("Modulo sysinfo cargado\n");
    return 0;
}

static void __exit sysinfo_exit(void)
{
    remove_proc_entry(PROC_NAME, NULL);
    pr_info("Modulo sysinfo descargado\n");
}

MODULE_LICENSE("GPL");
MODULE_AUTHOR("201700399");
MODULE_DESCRIPTION("Modulo de kernel para listar procesos del sistema");

module_init(sysinfo_init);
module_exit(sysinfo_exit);
