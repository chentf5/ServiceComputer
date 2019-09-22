package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

type selpg_args struct {
	start_page  int
	end_page    int
	in_filename string
	dest        string
	page_len    int
	page_type   int
}

var now_args selpg_args
var program_name string
var arg_n int

func Usage() {
	fmt.Println("\nUsage of selpg.")
	fmt.Println("\tselpg -s=Number -e=Number [options] [filename]")
	fmt.Println("\t-l:Determine the number of lines per page and default is 72.")
	fmt.Println("\t-f:Determine the type and the way to be seprated.")
	fmt.Println("\t-d:Determine the destination of output.")
	fmt.Println("\t[filename]: Read input from this file.")
	fmt.Println("\tIf filename is not given, read input from stdin. and Ctrl+D to cut out.")
}


func process_args(input_args []string)  {
	if len(input_args) < 3 {
		fmt.Fprintf(os.Stderr, "%s: not enough arguments\n", program_name)
		Usage()
		os.Exit(1)
	}

	if input_args[1][0] != '-' || input_args[1][1] != 's' {
		fmt.Fprintf(os.Stderr,"%s: 1st arg should be -sstart_page\n",program_name)
		Usage()
		os.Exit(1)
	}

	sp, _ := strconv.Atoi(input_args[1][2:])
	if sp < 1 {
		fmt.Fprintf(os.Stderr, "%s: invalid start page %d\n", program_name, sp)
		Usage()
		os.Exit(1)
	}
	now_args.start_page = sp

	if input_args[2][0] != '-' || input_args[2][1] != 'e' {
		fmt.Fprintf(os.Stderr, "%s: 2nd arg should be -eend_page\n", program_name)
		Usage()
		os.Exit(1)
	}

	ep, _ := strconv.Atoi(input_args[2][2:])
	if ep < 1 || ep < sp {
		fmt.Fprintf(os.Stderr, "%s: invalid end page %d\n", program_name, ep)
		Usage()
		os.Exit(1)
	}
	now_args.end_page = ep

	argindex := 3
	for {
		if argindex > arg_n - 1 || input_args[argindex][0] != '-'	{
			break
		}
		switch input_args[argindex][1] {
		case 'l':
			//获取一页的长度
			pl, _ := strconv.Atoi(input_args[argindex][2:])
			if pl < 1 {
				fmt.Fprintf(os.Stderr, "%s: invalid page length %d\n", program_name, pl)
				Usage()
				os.Exit(1)
			}
			now_args.page_len = pl
			argindex++
		case 'f':
			if len(input_args[argindex]) > 2 {
				fmt.Fprintf(os.Stderr, "%s: option should be \"-f\"\n", program_name)
				Usage()
				os.Exit(1)
			}
			now_args.page_type = 'f'
			argindex++
		case 'd':
			if len(input_args[argindex]) <= 2 {
				fmt.Fprintf(os.Stderr, "%s: -d option requires a printer destination\n", program_name)
				Usage()
				os.Exit(1)
			}
			now_args.dest = input_args[argindex][2:]
			argindex++
		default:
			fmt.Fprintf(os.Stderr, "%s: unknown option", program_name)
			Usage()
			os.Exit(1)
			
		}
	}
	if argindex <= arg_n-1 {
		now_args.in_filename = input_args[argindex]
	}
}

func process_input()	{
	var cmd *exec.Cmd
	var cmd_in io.WriteCloser
	var cmd_out io.ReadCloser
	if now_args.dest != "" {
		cmd = exec.Command("bash", "-c", now_args.dest)
		cmd_in, _ = cmd.StdinPipe()
		cmd_out, _ = cmd.StdoutPipe()
		//执行设定的命令
		cmd.Start()
	}

	if now_args.in_filename != "" {
		inf, err := os.Open(now_args.in_filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		line_count := 1
		page_count := 1
		fin := bufio.NewReader(inf)
		for {
			//读取输入文件中的一行数据
			line, _, err := fin.ReadLine()
			if err != io.EOF && err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			if err == io.EOF {
				break
			}
			if page_count >= now_args.start_page && page_count <= now_args.end_page {
				if now_args.dest == "" {
					//打印到屏幕
					fmt.Println(string(line))
				} else {
					//写入文件中
					fmt.Fprintln(cmd_in, string(line))
				}
			}
			line_count++
			if now_args.page_type == 'l' {
				if line_count > now_args.page_len {
					line_count = 1
					page_count++
				}
			} else {
				if string(line) == "\f" {
					page_count++
				}
			}
		}
		if now_args.dest != "" {
			cmd_in.Close()
			cmdBytes, err := ioutil.ReadAll(cmd_out)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Print(string(cmdBytes))
			//等待command退出
			cmd.Wait()
		}
	} else {
		//从标准输入读取内容
		ns := bufio.NewScanner(os.Stdin)
		line_count := 1
		page_count := 1
		out := ""

		for ns.Scan() {
			line := ns.Text()
			line += "\n"
			if page_count >= now_args.start_page && page_count <= now_args.end_page {
				out += line
			}
			line_count++
			if now_args.page_type == 'l' {
				if line_count > now_args.page_len {
					line_count = 1
					page_count++
				}
			} else {
				if string(line) == "\f" {
					page_count++
				}
			}
		}
		if now_args.dest == "" {
			fmt.Print(out)
		} else {
			fmt.Fprint(cmd_in, out)
			cmd_in.Close()
			cmdBytes, err := ioutil.ReadAll(cmd_out)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Print(string(cmdBytes))
			//等待command退出
			cmd.Wait()
		}
	}
}


func main() {
	input_args := os.Args
	now_args.start_page = 1
	now_args.end_page = 1
	now_args.in_filename = ""
	now_args.dest = ""
	now_args.page_len = 20 //默认20行一页
	now_args.page_type = 'l'
	arg_n = len(input_args)
	process_args(input_args)
	process_input()
}