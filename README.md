# xlsx2csv
用go写的一个将excel转换成csv的小工具，win/linux下均有效

###### 使用方法

xlsx2csv file_path [new_file_name]

其中filepath为你要转换的源文件，new_file_name为转换后的新文件名

当参数只有一个时，新文件和源文件在同一目录，新文件名为"原文件名_sheet名"

当参数为2个时，无论源文件有多少sheet，最后新文件将以最后一个sheet为其基准转化，当新文件名不包含路径的时候，新文件生成于源文件同一目录；当包含路径时，新文件保存于指定路径下