# GoSub: Submit PBS files 

This program helps submit PBS files in file/directory format one by one.
If an error is encountered, the program will re-submit in 5 minutes.

## Download

Go to [release page](https://github.com/ShixiangWang/gosub/releases) and select a proper version.

An example:

```bash
wget -c https://github.com/ShixiangWang/gosub/releases/download/v0.3.1/gosub_0.3.1_Linux_x86_64.tar.gz
tar zxvf gosub_0.3.1_Linux_x86_64.tar.gz
chmod u+x gosub
./gosub
```

## Usage

```bash
./gosub <path_to_PBS>
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

MIT