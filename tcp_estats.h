#ifndef TCP_ESTATS_H
#define TCP_ESTATS_H

#include <linux/types.h>

#define TCP_ESTATS_ACTIVE 1
#define TCP_ESTATS_INACTIVE 2

enum tcp_estats_sndlim_states {
  TCP_ESTATS_SNDLIM_NONE = -1,
  TCP_ESTATS_SNDLIM_SENDER,
  TCP_ESTATS_SNDLIM_CWND,
  TCP_ESTATS_SNDLIM_RWIN,
  TCP_ESTATS_SNDLIM_STARTUP,
  TCP_ESTATS_SNDLIM_TSODEFER,
  TCP_ESTATS_SNDLIM_PACE,
  TCP_ESTATS_SNDLIM_NSTATES
};

enum tcp_estats_addrtype {
  TCP_ESTATS_ADDRTYPE_IPV4 = 1,
  TCP_ESTATS_ADDRTYPE_IPV6 = 2
};

enum tcp_estats_operation {
  TCP_ESTATS_OPERATION_SET,
  TCP_ESTATS_OPERATION_ADD,
  TCP_ESTATS_OPERATION_SUB,
  TCP_ESTATS_OPERATION_MAX,
  TCP_ESTATS_OPERATION_MIN,
};

enum tcp_estats_global_table {
  TCP_ESTATS_GLOBAL_TABLE_LIMSTATE,
  TCP_ESTATS_GLOBAL_TABLE_LIMSTATE_TS,
  TCP_ESTATS_GLOBAL_TABLE_START_TS,
  TCP_ESTATS_GLOBAL_TABLE_CURRENT_TS,
  TCP_ESTATS_GLOBAL_TABLE_START_TV
};

enum tcp_estats_connection_table {
  TCP_ESTATS_CONNECTION_TABLE_ADDRESS_TYPE,
  TCP_ESTATS_CONNECTION_TABLE_LOCAL_ADDRESS,
  TCP_ESTATS_CONNECTION_TABLE_REMOTE_ADDRESS,
  TCP_ESTATS_CONNECTION_TABLE_LOCAL_PORT,
  TCP_ESTATS_CONNECTION_TABLE_REMOTE_PORT
};

enum tcp_estats_perf_table {
  TCP_ESTATS_PERF_TABLE_SEGSOUT,
  TCP_ESTATS_PERF_TABLE_DATASEGSOUT,
  TCP_ESTATS_PERF_TABLE_DATAOCTETSOUT,  // u64
  TCP_ESTATS_PERF_TABLE_SEGSRETRANS,
  TCP_ESTATS_PERF_TABLE_OCTETSRETRANS,
  TCP_ESTATS_PERF_TABLE_SEGSIN,
  TCP_ESTATS_PERF_TABLE_DATASEGSIN,
  TCP_ESTATS_PERF_TABLE_DATAOCTETSIN,
  TCP_ESTATS_PERF_TABLE_NORMALRTTM,
  TCP_ESTATS_PERF_TABLE_HIGHRTTM,
  /*		ElapsedSecs */
  /*		ElapsedMicroSecs */
  /*		StartTimeStamp */
  /*		CurMSS */
  /*		PipeSize */
  TCP_ESTATS_PERF_TABLE_MAXPIPESIZE,
  /*		SmoothedRTT */
  /*		CurRTO */
  TCP_ESTATS_PERF_TABLE_CONGSIGNALS,
  /*		CurCwnd */
  /*		CurSsthresh */
  TCP_ESTATS_PERF_TABLE_TIMEOUTS,
  /*		CurRwinSent */
  TCP_ESTATS_PERF_TABLE_MAXRWINSENT,
  TCP_ESTATS_PERF_TABLE_ZERORWINSENT,
  /*		CurRwinRcvd */
  TCP_ESTATS_PERF_TABLE_MAXRWINRCVD,
  TCP_ESTATS_PERF_TABLE_ZERORWINRCVD,
  /*		SndLimTransRwin */
  /*		SndLimTransCwnd */
  /*		SndLimTransSnd */
  /*		SndLimTimeRwin */
  /*		SndLimTimeCwnd */
  /*		SndLimTimeSnd */
  // TODO: figure out how to do this
  // u32		snd_lim_trans[TCP_ESTATS_SNDLIM_NSTATES];
  // u32		snd_lim_time[TCP_ESTATS_SNDLIM_NSTATES];
};

enum tcp_estats_path_table {
  TCP_ESTATS_PATH_TABLE_NONRECOVDAEPISODES,
  TCP_ESTATS_PATH_TABLE_SUMOCTETSREORDERED,
  TCP_ESTATS_PATH_TABLE_NONRECOVDA,
  TCP_ESTATS_PATH_TABLE_SAMPLERTT,
  TCP_ESTATS_PATH_TABLE_MAXRTT,
  TCP_ESTATS_PATH_TABLE_MINRTT,
  TCP_ESTATS_PATH_TABLE_SUMRTT,
  TCP_ESTATS_PATH_TABLE_COUNTRTT,
  TCP_ESTATS_PATH_TABLE_MAXRTO,
  TCP_ESTATS_PATH_TABLE_MINRTO,
  TCP_ESTATS_PATH_TABLE_IPTTL,    // u8
  TCP_ESTATS_PATH_TABLE_IPTOSIN,  // u8
  TCP_ESTATS_PATH_TABLE_PRECONGSUMCWND,
  TCP_ESTATS_PATH_TABLE_PRECONGSUMRTT,
  TCP_ESTATS_PATH_TABLE_POSTCONGSUMRTT,
  TCP_ESTATS_PATH_TABLE_POSTCONGCOUNTRTT,
  TCP_ESTATS_PATH_TABLE_ECNSIGNALS,
  TCP_ESTATS_PATH_TABLE_DUPACKEPISODES,
  TCP_ESTATS_PATH_TABLE_DUPACKSOUT,
  TCP_ESTATS_PATH_TABLE_CERCVD,
  TCP_ESTATS_PATH_TABLE_ECESENT
};

enum tcp_estats_stack_table {
  TCP_ESTATS_STACK_TABLE_ACTIVEOPEN,
  TCP_ESTATS_STACK_TABLE_MAXSSCWND,
  TCP_ESTATS_STACK_TABLE_MAXCACWND,
  TCP_ESTATS_STACK_TABLE_MAXSSTHRESH,
  TCP_ESTATS_STACK_TABLE_MINSSTHRESH,
  TCP_ESTATS_STACK_TABLE_DUPACKSIN,
  TCP_ESTATS_STACK_TABLE_SPURIOUSFRDETECTED,
  TCP_ESTATS_STACK_TABLE_SPURIOUSRTODETECTED,
  TCP_ESTATS_STACK_TABLE_SOFTERRORS,
  TCP_ESTATS_STACK_TABLE_SOFTERRORREASON,
  TCP_ESTATS_STACK_TABLE_SLOWSTART,
  TCP_ESTATS_STACK_TABLE_CONGAVOID,
  TCP_ESTATS_STACK_TABLE_OTHERREDUCTIONS,
  TCP_ESTATS_STACK_TABLE_CONGOVERCOUNT,
  TCP_ESTATS_STACK_TABLE_FASTRETRAN,
  TCP_ESTATS_STACK_TABLE_SUBSEQUENTTIMEOUTS,
  TCP_ESTATS_STACK_TABLE_ABRUPTTIMEOUTS,
  TCP_ESTATS_STACK_TABLE_SACKSRCVD,
  TCP_ESTATS_STACK_TABLE_SACKBLOCKSRCVD,
  TCP_ESTATS_STACK_TABLE_SENDSTALL,
  TCP_ESTATS_STACK_TABLE_DSACKDUPS,
  TCP_ESTATS_STACK_TABLE_MAXMSS,
  TCP_ESTATS_STACK_TABLE_MINMSS,
  TCP_ESTATS_STACK_TABLE_SNDINITIAL,
  TCP_ESTATS_STACK_TABLE_RECINITIAL,
  TCP_ESTATS_STACK_TABLE_CURRETXQUEUE,
  TCP_ESTATS_STACK_TABLE_MAXRETXQUEUE,
  TCP_ESTATS_STACK_TABLE_MAXREASMQUEUE,
  TCP_ESTATS_STACK_TABLE_EARLYRETRANS,
  TCP_ESTATS_STACK_TABLE_EARLYRETRANSDELAY
};

enum tcp_estats_app_table {
  TCP_ESTATS_APP_TABLE_SNDMAX,
  TCP_ESTATS_APP_TABLE_THRUOCTETSACKED,
  // TODO: more
};

enum tcp_estats_extras_table {
  TCP_ESTATS_EXTRAS_TABLE_OTHERREDUCTIONSCV,
  TCP_ESTATS_EXTRAS_TABLE_OTHERREDUCTIONSCM,
  TCP_ESTATS_EXTRAS_TABLE_PRIORITY
};

// An entry describes the operation to make on one of the estats tables.
struct key {
  // TODO: ipv6
  __u64 pid_tgid;
  __u32 saddr;
  __u32 daddr;
  __u16 sport;
  __u16 dport;
};

struct entry {
  struct key key;
  enum tcp_estats_operation op;
  // oneof table var enum entries
  __u32 var;
  __u32 val;
};

#endif  // TCP_ESTATS_H
