#include <linux/build-salt.h>
#include <linux/module.h>
#include <linux/vermagic.h>
#include <linux/compiler.h>

BUILD_SALT;

MODULE_INFO(vermagic, VERMAGIC_STRING);
MODULE_INFO(name, KBUILD_MODNAME);

__visible struct module __this_module
__section(.gnu.linkonce.this_module) = {
	.name = KBUILD_MODNAME,
	.init = init_module,
#ifdef CONFIG_MODULE_UNLOAD
	.exit = cleanup_module,
#endif
	.arch = MODULE_ARCH_INIT,
};

#ifdef CONFIG_RETPOLINE
MODULE_INFO(retpoline, "Y");
#endif

static const struct modversion_info ____versions[]
__used __section(__versions) = {
	{ 0xe4c970fb, "module_layout" },
	{ 0x50049ec8, "single_release" },
	{ 0xc353a5e3, "seq_read" },
	{ 0xddfb6d9e, "seq_lseek" },
	{ 0x1a1279de, "remove_proc_entry" },
	{ 0xc5850110, "printk" },
	{ 0xf914135c, "proc_create" },
	{ 0xdecd0b29, "__stack_chk_fail" },
	{ 0xb978c7fb, "init_task" },
	{ 0xd1639f17, "seq_printf" },
	{ 0xd2b23511, "seq_puts" },
	{ 0x40c7247c, "si_meminfo" },
	{ 0x19867527, "single_open" },
	{ 0xbdfb6dbb, "__fentry__" },
};

MODULE_INFO(depends, "");


MODULE_INFO(srcversion, "5B09E3ADE41021023B408A7");
