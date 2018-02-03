package main
 import(
	 "errors"
	 "fmt"
 	"os"

	 "github.com/urfave/cli"
 )

 func transAction(file_name, new_name string) error {
	 var err = fileCheck(file_name)
	 if (err != nil) {
		 return err
	 }
	 transXlsx2Csv(file_name, new_name)
	 return nil
 }

 func argCheck(c *cli.Context) error {
	 var arg_num = c.NArg()
	 if arg_num <= 0 {
	 	var msg = fmt.Sprintf("arg num error, %s", c.App.UsageText)
	 	return errors.New(msg)
	 }
	 return nil
 }

 func main() {
 	var app = cli.NewApp()
 	app.Name = "xlsx2csv"
 	app.Usage = "a command tool trans excel to csv"
	app.Version = "1.0.0"
	app.UsageText = fmt.Sprintf("please %s file_path [new_file_name]", app.Name)
	app.ArgsUsage = "path of trans file"

	app.Action = func(c *cli.Context) error {
		//var arg_num = c.NArg()
		var err = argCheck(c)
		if err != nil {
			fmt.Println(err)
			return err
		}
		var file_name = c.Args().Get(0)
		var new_name = c.Args().Get(1)

		err = transAction(file_name, new_name)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	}

	app.Run(os.Args)
 }