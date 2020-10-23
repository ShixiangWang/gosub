# GoSub: Automatically Submit PBS Files to Free Your Hands

This program helps submit PBS files in file/directory format one by one or generate a work pbs to parallelly run all PBS files.

If an error is encountered, the program will re-submit in 5 minutes.

## Download and Install

Pick the latest version of `gosub` and install it, e.g.:

```bash
wget -c https://github.com/ShixiangWang/gosub/releases/download/v1.3/gosub_1.3_Linux_x86_64.tar.gz
tar zxvf gosub_1.3_Linux_x86_64.tar.gz
chmod u+x gosub
./gosub -h
```

## Usage

### Simple mode

```bash
./gosub <path_to_PBS>
```

> Both file path and directory path are supported.

A better way is to use it with nohup and `&`.

```bash
nohup ./gosub <path_to_PBS> &
```

So you can free your terminal. Just remember to record the PID in case you want to kill it.

### Parallel mode

In parallel mode, `gosub` will generate a new work PBS file and use [`rush`](https://github.com/shenwei356/rush) to parallelly run your all PBS files.

```sh
./gosub -h
Usage of ./gosub:
  -hold
        set it if you want to check and qsub by your own. Only work when -p enabled.
  -jobs int
        run n jobs in parallel, at default will use nodes*ppn. Only work when -p enabled.
  -mem string
        memory size, e.g. 5gb. Only work when -p enabled. (default "auto")
  -name string
        an file prefix for generating output PBS file. Only work when -p enabled. (default "pwork")
  -nodes int
        an int to specify node number to use. Only work when -p enabled. (default 1)
  -p    enable parallel processing.
  -ppn int
        an int to specify cpu number per node. Only work when -p enabled. (default 1)
  -walltime string
        walltime setting. Only work when -p enabled. (default "24:00:00")
```

For example:

```sh
./gosub -p -nodes 3 -ppn 5 -mem 5gb testp
```

## Tests

I tested it in a HPC account.

```bash
$ biosoft/gosub test_gosub/
gosub version: 0.3.1
Starting...
Submitted file list will be
  save to success_submitted_list.txt
====================================
2020/03/30 22:20:35 Submitting test_gosub/test.pbs
2020/03/30 22:20:35 Submitting test_gosub/test_1.pbs
2020/03/30 22:20:35 Submitting test_gosub/test_10.pbs
2020/03/30 22:20:35 Submitting test_gosub/test_100.pbs
2020/03/30 22:20:35 Submitting test_gosub/test_100.pbs failed with error:
2020/03/30 22:20:35 exit status 200
2020/03/30 22:20:35 Waiting for 5 minutes..
2020/03/30 22:25:35 Calling back to submit...
2020/03/30 22:25:35 Submitting test_gosub/test_100.pbs
2020/03/30 22:25:35 Submitting test_gosub/test_11.pbs
2020/03/30 22:25:35 Submitting test_gosub/test_12.pbs
2020/03/30 22:25:35 Submitting test_gosub/test_13.pbs
2020/03/30 22:25:35 Submitting test_gosub/test_14.pbs
2020/03/30 22:25:35 Submitting test_gosub/test_15.pbs
2020/03/30 22:25:35 Submitting test_gosub/test_16.pbs
2020/03/30 22:25:35 Submitting test_gosub/test_17.pbs
2020/03/30 22:25:35 Submitting test_gosub/test_18.pbs
2020/03/30 22:25:35 Submitting test_gosub/test_19.pbs
2020/03/30 22:25:35 Submitting test_gosub/test_2.pbs
2020/03/30 22:25:35 Submitting test_gosub/test_20.pbs
2020/03/30 22:25:35 Submitting test_gosub/test_21.pbs
2020/03/30 22:25:35 Submitting test_gosub/test_22.pbs
2020/03/30 22:25:35 Submitting test_gosub/test_23.pbs
2020/03/30 22:25:35 Submitting test_gosub/test_24.pbs
2020/03/30 22:25:35 Submitting test_gosub/test_24.pbs failed with error:
2020/03/30 22:25:35 exit status 200
2020/03/30 22:25:35 Waiting for 5 minutes..
2020/03/30 22:30:36 Calling back to submit...
2020/03/30 22:30:36 Submitting test_gosub/test_24.pbs
2020/03/30 22:30:36 Submitting test_gosub/test_25.pbs
2020/03/30 22:30:36 Submitting test_gosub/test_26.pbs
2020/03/30 22:30:36 Submitting test_gosub/test_27.pbs
2020/03/30 22:30:36 Submitting test_gosub/test_28.pbs
2020/03/30 22:30:36 Submitting test_gosub/test_29.pbs
2020/03/30 22:30:36 Submitting test_gosub/test_3.pbs
2020/03/30 22:30:36 Submitting test_gosub/test_30.pbs
2020/03/30 22:30:36 Submitting test_gosub/test_31.pbs
2020/03/30 22:30:36 Submitting test_gosub/test_32.pbs
2020/03/30 22:30:36 Submitting test_gosub/test_33.pbs
2020/03/30 22:30:36 Submitting test_gosub/test_34.pbs
2020/03/30 22:30:36 Submitting test_gosub/test_35.pbs
2020/03/30 22:30:36 Submitting test_gosub/test_36.pbs
2020/03/30 22:30:36 Submitting test_gosub/test_37.pbs
2020/03/30 22:30:37 Submitting test_gosub/test_38.pbs
2020/03/30 22:30:37 Submitting test_gosub/test_39.pbs
2020/03/30 22:30:37 Submitting test_gosub/test_39.pbs failed with error:
2020/03/30 22:30:37 exit status 200
2020/03/30 22:30:37 Waiting for 5 minutes..
```

## LICENSE

MIT@ShixiangWang, 2020.
