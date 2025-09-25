#include <linux/init.h>
#include <linux/module.h>
#include <linux/proc_fs.h>
#include <linux/seq_file.h>
#include <linux/sched/signal.h>
#include <linux/mm.h>

#define PROC_NAME "continfo_so1_201700399"  // tu carnet

// Función que genera la salida en /proc
static int proc_show(struct seq_file *m, void *v) {
    struct task_struct *task;

    seq_puts(m, "{\n  \"processes\": [\n");

    for_each_process(task) {
        if (task->mm) {
            unsigned long vsize_kb = (task->mm->total_vm * PAGE_SIZE) / 1024;
            unsigned long rss_kb   = (get_mm_rss(task->mm) * PAGE_SIZE) / 1024;

            seq_printf(m,
                "    {\"pid\": %d, \"comm\": \"%s\", \"vsz_kb\": %lu, \"rss_kb\": %lu},\n",
                task->pid, task->comm, vsize_kb, rss_kb
            );
        }
    }

    seq_puts(m, "    {}\n  ]\n}\n");
    return 0;
}

// Abrir el archivo /proc
static int proc_open(struct inode *inode, struct file *file) {
    return single_open(file, proc_show, NULL);
}

// Kernel 5.4 usa struct file_operations
static const struct file_operations proc_fops = {
    .owner   = THIS_MODULE,
    .open    = proc_open,
    .read    = seq_read,
    .llseek  = seq_lseek,
    .release = single_release,
};


static int __init continfo_init(void) {
    if (!proc_create(PROC_NAME, 0444, NULL, &proc_fops)) {
        pr_err("continfo: error creando /proc/%s\n", PROC_NAME);
        return -ENOMEM;
    }
    pr_info("continfo: módulo cargado correctamente\n");
    return 0;
}

static void __exit continfo_exit(void) {
    remove_proc_entry(PROC_NAME, NULL);
    pr_info("continfo: módulo descargado\n");
}

module_init(continfo_init);
module_exit(continfo_exit);

MODULE_LICENSE("GPL");
MODULE_AUTHOR("201700399");
MODULE_DESCRIPTION("Modulo SO1 que lista procesos en JSON");
