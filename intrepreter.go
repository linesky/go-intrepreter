package main

import ("fmt"
"io/ioutil"
"os"
"os/exec"
"strings"
)
func srun(datas string){
	arg:=strings.Split(datas," ");
	a:=strings.Count(datas," ");
	if a==0{
		cmd,err:=exec.Command(arg[0]).Output();
		if err != nil{
			fmt.Println(err);
		}
		fmt.Println(string(cmd));
	}
	if a==1{
		cmd,err:=exec.Command(arg[0],arg[1]).Output();
		if err != nil{
			fmt.Println(err);
		}

		fmt.Println(string(cmd));
	}
	if a==2{
		cmd,err:=exec.Command(arg[0],arg[1],arg[2]).Output();
		if err != nil{
			fmt.Println(err);
		}

		fmt.Println(string(cmd));
	}
	if a==3{
		cmd,err:=exec.Command(arg[0],arg[1],arg[2],arg[3]).Output();
		if err != nil{
			fmt.Println(err);
		}

		fmt.Println(string(cmd));
	}
	if a==4{
		cmd,err:=exec.Command(arg[0],arg[1],arg[2],arg[3],arg[4]).Output();
		if err != nil{
			fmt.Println(err);
		}

		fmt.Println(string(cmd));
	}
	if a==5{
		cmd,err:=exec.Command(arg[0],arg[1],arg[2],arg[3],arg[4],arg[5]).Output();
		if err != nil{
			fmt.Println(err);
		}

		fmt.Println(string(cmd));
	}
	if a==6{
		cmd,err:=exec.Command(arg[0],arg[1],arg[2],arg[3],arg[4],arg[5],arg[6]).Output();
		if err != nil{
			fmt.Println(err);
		}

		fmt.Println(string(cmd));
	}
	if a==7{
		cmd,err:=exec.Command(arg[0],arg[1],arg[2],arg[3],arg[4],arg[5],arg[6],arg[7]).Output();
		if err != nil{
			fmt.Println(err);
		}

		fmt.Println(string(cmd));
	}
	if a>7{
		cmd,err:=exec.Command(arg[0],arg[1],arg[2],arg[3],arg[4],arg[5],arg[6],arg[7],arg[8]).Output();
		if err != nil{
			fmt.Println(err);
		}
		fmt.Println(string(cmd));
	}

}
func scommand(datas string){
	ss:=strings.Split(datas,";");
	for a:= range ss{
		if ss[a] != "" {
			srun(ss[a]);
		}
	}
}
func sline(datas string){
	ddatas,err:= ioutil.ReadFile(datas);
	if err != nil {
		fmt.Println(err);
	}
	s:=string(ddatas);
	ss:=strings.Split(s,"\n");
	for a:= range ss{
		if ss[a] != "" {
			scommand(ss[a]);
		}
	}
	
}

func main(){
	arg :=os.Args[1:]
	sline(arg[0]);
}

