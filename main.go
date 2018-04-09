package main
 import(
	 "errors"
	 "fmt"
 	"os"

	 "github.com/urfave/cli"
	 "xlsx2csv/config"
 )

 func transAction(file_name, new_name string) error {
	 var file_type, err = getFileType(file_name)
	 if err != nil {
	 	return err
	 }

	 err = nil
	 switch file_type {
	 case config.FILE_TYPE_FILE:
	 	err = transFile(file_name, new_name)
	 	break
	 case config.FILE_TYPE_DIR:
	 	err = transDir(file_name, new_name)
	 	break
	 }

	 return err
 }



 func argCheck(c *cli.Context) error {
	 var arg_num = c.NArg()
	 if arg_num <= 0 {
	 	var msg = fmt.Sprintf("arg num error, %s", c.App.UsageText)
	 	return errors.New(msg)
	 }
	 return nil
 }

 func transParams(client *cli.Context) (file_name, new_name string) {
	var src_name = client.Args().Get(0)
	var arg_num = client.NArg()
	var dst_name = ""
	if (arg_num == 2) {
		dst_name = client.Args().Get(1)
	}
	return src_name, dst_name
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
		file_name, new_name := transParams(c)

		err = transAction(file_name, new_name)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	}

	app.Run(os.Args)
 }