package main

import "fmt"

type Student struct {
	Chinese int
	Math    int
	English int
}

var ss = []Student{
	{80, 80, 80},
	{70, 80, 80},
	{70, 80, 60},
	{90, 80, 59},
	{90, 40, 59},
	{190, 40, 59},
	{80, 75, 66},
}

func main()  {
	//runError()
	runPanicRecover()
}

func runError() {
	i := 0
	for ; i < len(ss); i++ {
		flag, err := checkStudent(&ss[i])
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("student %#v,及格? ：%t \n", ss[i], flag)
	}
}

func runPanicRecover() {
	i := 0
	defer func() {

		if err := recover(); err != nil {
			fmt.Println(err)
		}
		i++
		for ; i < len(ss); i++ {
			fmt.Printf("student %#v,及格? ：%t \n", ss[i], checkStudentS(&ss[i]))
		}

	}()

	for ; i < len(ss); i++ {
		fmt.Printf("student %#v,及格? ：%t \n", ss[i], checkStudentS(&ss[i]))
	}
}

func checkStudent(s *Student) (bool, error) {
	if s.Chinese > 100 || s.Math > 100 || s.English > 100 {
		return false, fmt.Errorf("student %#v, something error", s)
	}

	if s.Chinese > 60 && s.Math > 60 && s.English > 60 {
		return true, nil
	}

	return false, nil
}

func checkStudentS(s *Student) bool {
	if s.Chinese > 100 || s.Math > 100 || s.English > 100 {
		panic(fmt.Errorf("student %#v, something error", s))
	}

	if s.Chinese > 60 && s.Math > 60 && s.English > 60 {
		return true
	}

	return false
}
