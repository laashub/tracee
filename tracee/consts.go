package tracee

// Syscalls is a list of system calls that tracee can trace
var Syscalls = []string{"execve", "execveat", "mmap", "mprotect", "clone", 
"fork", "vfork", "newstat", 	"newfstat", "newlstat", "mknod", "mknodat", 
"dup", "dup2", "dup3", 	"memfd_create", "socket", "close", "ioctl", "access", 
"faccessat", "kill", "listen", 	"connect", "accept", "accept4", "bind", 
"getsockname", "prctl", "ptrace", 	"process_vm_writev", "process_vm_readv", 
"init_module", "finit_module", "delete_module", 	"symlink", "symlinkat", "getdents", 
"getdents64", "creat", "open", "openat", 	"mount", "umount", "unlink", "unlinkat", 
"setuid", "setgid", "setreuid", "setregid", 	"setresuid", "setresgid", "setfsuid", 
"setfsgid"}
// Sysevents is a list of system events that tracee can trace
var Sysevents = []string{"cap_capable", "do_exit"}
// essentialSyscalls is a list of system calls that are essential to the operation of tracee and therefore must be traced
var essentialSyscalls = []string{"execve", "execveat"}
// essentialSysevents is a list of system events that are essential to the operation of tracee and therefore must be traced
var essentialSysevents = []string{"do_exit"}

// ArgType is an enum that encodes the argument types that the BPF program may write to the shared buffer
type ArgType uint8
const (
	NONE          ArgType = 0
	INT_T         ArgType = 1
	UINT_T        ArgType = 2
	LONG_T        ArgType = 3
	ULONG_T       ArgType = 4
	OFF_T_T       ArgType = 5
	MODE_T_T      ArgType = 6
	DEV_T_T       ArgType = 7
	SIZE_T_T      ArgType = 8
	POINTER_T     ArgType = 9
	STR_T         ArgType = 10
	STR_ARR_T     ArgType = 11
	SOCKADDR_T    ArgType = 12
	OPEN_FLAGS_T  ArgType = 13
	EXEC_FLAGS_T  ArgType = 14
	SOCK_DOM_T    ArgType = 15
	SOCK_TYPE_T   ArgType = 16
	CAP_T         ArgType = 17
	SYSCALL_T     ArgType = 18
	PROT_FLAGS_T  ArgType = 19
	ACCESS_MODE_T ArgType = 20
	TYPE_MAX      ArgType = 255
)

// This array maps event id in `context.event_id` to event name
// The id is the index in the array (ids are serial)  
// TODO: make this a map? because index is semantic
var eventNames = []string{
"read", "write", "open", "close", "stat", "fstat", "lstat", "poll", "lseek", 
"mmap", "mprotect", "munmap", "brk", "rt_sigaction", "rt_sigprocmask", 
"rt_sigreturn", "ioctl", "pread64", "pwrite64", "readv", "writev", "access", 
"pipe", "select", "sched_yield", "mremap", "msync", "mincore", "madvise", 
"shmget", "shmat", "shmctl", "dup", "dup2", "pause", "nanosleep", "getitimer", 
"alarm", "setitimer", "getpid", "sendfile", "socket", "connect", "accept", 
"sendto", "recvfrom", "sendmsg", "recvmsg", "shutdown", "bind", "listen", 
"getsockname", "getpeername", "socketpair", "setsockopt", "getsockopt", 
"clone", "fork", "vfork", "execve", "exit", "wait4", "kill", "uname", "semget",
"semop", "semctl", "shmdt", "msgget", "msgsnd", "msgrcv", "msgctl", "fcntl", 
"flock", "fsync", "fdatasync", "truncate", "ftruncate", "getdents", "getcwd", 
"chdir", "fchdir", "rename", "mkdir", "rmdir", "creat", "link", "unlink", 
"symlink", "readlink", "chmod", "fchmod", "chown", "fchown", "lchown", "umask",
 "gettimeofday", "getrlimit", "getrusage", "sysinfo", "times", "ptrace", 
 "getuid", "syslog", "getgid", "setuid", "setgid", "geteuid", "getegid", 
 "setpgid", "getppid", "getpgrp", "setsid", "setreuid", "setregid", 
 "getgroups", "setgroups", "setresuid", "getresuid", "setresgid", "getresgid", 
 "getpgid", "setfsuid", "setfsgid", "getsid", "capget", "capset", 
 "rt_sigpending", "rt_sigtimedwait", "rt_sigqueueinfo", "rt_sigsuspend", 
 "sigaltstack", "utime", "mknod", "uselib", "personality", "ustat", "statfs", 
 "fstatfs", "sysfs", "getpriority", "setpriority", "sched_setparam", 
 "sched_getparam", "sched_setscheduler", "sched_getscheduler", 
 "sched_get_priority_max", "sched_get_priority_min", "sched_rr_get_interval", 
 "mlock", "munlock", "mlockall", "munlockall", "vhangup", "modify_ldt", 
 "pivot_root", "sysctl", "prctl", "arch_prctl", "adjtimex", "setrlimit", 
 "chroot", "sync", "acct", "settimeofday", "mount", "umount", "swapon", 
 "swapoff", "reboot", "sethostname", "setdomainname", "iopl", "ioperm", 
 "create_module", "init_module", "delete_module", "get_kernel_syms", 
 "query_module", "quotactl", "nfsservctl", "getpmsg", "putpmsg", "afs", 
 "tuxcall", "security", "gettid", "readahead", "setxattr", "lsetxattr", 
 "fsetxattr", "getxattr", "lgetxattr", "fgetxattr", "listxattr", "llistxattr", 
 "flistxattr", "removexattr", "lremovexattr", "fremovexattr", "tkill", "time", 
 "futex", "sched_setaffinity", "sched_getaffinity", "set_thread_area", 
 "io_setup", "io_destroy", "io_getevents", "io_submit", "io_cancel", 
 "get_thread_area", "lookup_dcookie", "epoll_create", "epoll_ctl_old", 
 "epoll_wait_old", "remap_file_pages", "getdents64", "set_tid_address", 
 "restart_syscall", "semtimedop", "fadvise64", "timer_create", "timer_settime",
	"timer_gettime", "timer_getoverrun", "timer_delete", "clock_settime", 
	"clock_gettime", "clock_getres", "clock_nanosleep", "exit_group", 
	"epoll_wait", "epoll_ctl", "tgkill", "utimes", "vserver", "mbind", 
	"set_mempolicy", "get_mempolicy", "mq_open", "mq_unlink", "mq_timedsend", 
	"mq_timedreceive", "mq_notify", "mq_getsetattr", "kexec_load", "waitid", 
	"add_key", "request_key", "keyctl", "ioprio_set", "ioprio_get", 
	"inotify_init", "inotify_add_watch", "inotify_rm_watch", "migrate_pages", 
	"openat", "mkdirat", "mknodat", "fchownat", "futimesat", "newfstatat", 
	"unlinkat", "renameat", "linkat", "symlinkat", "readlinkat", "fchmodat", 
	"faccessat", "pselect6", "ppoll", "unshare", "set_robust_list", 
	"get_robust_list", "splice", "tee", "sync_file_range", "vmsplice", 
	"move_pages", "utimensat", "epoll_pwait", "signalfd", "timerfd_create", 
	"eventfd", "fallocate", "timerfd_settime", "timerfd_gettime", "accept4", 
	"signalfd4", "eventfd2", "epoll_create1", "dup3", "pipe2", "ionotify_init1", 
	"preadv", "pwritev", "rt_tgsigqueueinfo", "perf_event_open", "recvmmsg", 
	"fanotify_init", "fanotify_mark", "prlimit64", "name_tohandle_at", 
	"open_by_handle_at", "clock_adjtime", "sycnfs", "sendmmsg", "setns", 
	"getcpu", "process_vm_readv", "process_vm_writev", "kcmp", "finit_module", 
	"sched_setattr", "sched_getattr", "renameat2", "seccomp", "getrandom", 
	"memfd_create", "kexec_file_load", "bpf", "execveat", "userfaultfd", 
	"membarrier", "mlock2", "copy_file_range", "preadv2", "pwritev2", 
	"pkey_mprotect", "pkey_alloc", "pkey_free", "statx", "io_pgetevents", "rseq",
	// syscall events start here
	"do_exit", "cap_capable", 
}

type bpfConfig int
const (
	CONFIG_CONT_MODE    bpfConfig = 0
	CONFIG_SHOW_SYSCALL bpfConfig = 1
)