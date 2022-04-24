# ref: https://docs.google.com/presentation/d/1TPUEiLOx_T2_8sZNWSk5t_zg4ZYB4YB2bYYtl4vSB78/edit#slide=id.gca6063323d_3_77

strace-go:
	go build -o hello.out hello.go && strace ./hello.out | grep Hello

strace-py:
	strace python3.9 hello.py | grep Hello
