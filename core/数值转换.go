package E

//到数值
//数值到大写
//数值到金额
//数值到格式文本
//取十六进制文本
//取八进制文本
//十六进制
//二进制

//调用格式： 〈文本型〉 数值到大写 （双精度小数型 欲转换形式的数值，逻辑型 是否转换为简体） - 系统核心支持库->数值转换
//英文名称：UNum
//将数值转换为简体或繁体的大写形式，返回转换后的文本。本命令为初级命令。
//参数<1>的名称为“欲转换形式的数值”，类型为“双精度小数型（double）”。
//参数<2>的名称为“是否转换为简体”，类型为“逻辑型（bool）”，初始值为“假”。如果参数值为假，则转换为繁体。
//
//操作系统需求： Windows、Linux

//
//调用格式： 〈文本型〉 数值到金额 （双精度小数型 欲转换形式的数值，逻辑型 是否转换为简体） - 系统核心支持库->数值转换
//英文名称：NumToRMB
//将数值转换为金额的简体或繁体大写形式，返回转换后的文本。本命令为初级命令。
//参数<1>的名称为“欲转换形式的数值”，类型为“双精度小数型（double）”。
//参数<2>的名称为“是否转换为简体”，类型为“逻辑型（bool）”，初始值为“假”。如果参数值为假，则转换为繁体。
//
//操作系统需求： Windows、Linux

//调用格式： 〈文本型〉 数值到格式文本 （双精度小数型 欲转换为文本的数值，［整数型 小数保留位数］，逻辑型 是否进行千分位分隔） - 系统核心支持库->数值转换
//英文名称：NumToText
//返回一个文本，代表指定数值被格式转换后的结果。本命令为初级命令。
//参数<1>的名称为“欲转换为文本的数值”，类型为“双精度小数型（double）”。
//参数<2>的名称为“小数保留位数”，类型为“整数型（int）”，可以被省略。如果大于0，表示小数点右边应四舍五入保留的位数；如果等于0，表示舍入到整数；如果小于0，表示小数点左边舍入到的位置。例如：数值到格式文本 (1056.65, 1, 假) 返回 “1056.7”； 数值到格式文本 (1056.65, 0, 假) 返回 “1057”； 数值到格式文本 (1056.65, -1, 假) 返回 “1060”。如果省略本参数，则默认为保留所有实际存在的小数位。
//参数<3>的名称为“是否进行千分位分隔”，类型为“逻辑型（bool）”，初始值为“假”。如果参数值为真，数值文本的每个千分位上都将被自动插入一个逗号进行分隔。
//
//操作系统需求： Windows、Linux

//调用格式： 〈文本型〉 取十六进制文本 （整数型 欲取进制文本的数值） - 系统核心支持库->数值转换
//英文名称：GetHexText
//返回一个文本，代表指定数值的十六进制形式。本命令为初级命令。
//参数<1>的名称为“欲取进制文本的数值”，类型为“整数型（int）”。
//
//操作系统需求： Windows、Linux

//调用格式： 〈文本型〉 取八进制文本 （整数型 欲取进制文本的数值） - 系统核心支持库->数值转换
//英文名称：GetOctText
//返回一个文本，代表指定数值的八进制形式。本命令为初级命令。
//参数<1>的名称为“欲取进制文本的数值”，类型为“整数型（int）”。
//
//操作系统需求： Windows、Linux

//调用格式： 〈整数型〉 十六进制 （文本型 十六进制文本常量） - 系统核心支持库->数值转换
//英文名称：hex
//计算返回所指定十六进制文本常量对应的整数值. 本命令在编译时被直接预处理为整数型参量值,不影响程序执行效率.本命令为初级命令。
//参数<1>的名称为“十六进制文本常量”，类型为“文本型（text）”。本参数指定欲转换为整数的十六进制文本常量。
//
//操作系统需求： Windows、Linux

//调用格式： 〈整数型〉 二进制 （文本型 二进制文本常量） - 系统核心支持库->数值转换
//英文名称：binary
//计算返回所指定二进制文本常量对应的整数值. 本命令在编译时被直接预处理为整数型参量值,不影响程序执行效率.本命令为初级命令。
//参数<1>的名称为“二进制文本常量”，类型为“文本型（text）”。本参数指定欲转换为整数的二进制文本常量。
//
//操作系统需求： Windows、Linux
