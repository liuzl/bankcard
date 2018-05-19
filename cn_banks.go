package bankcard

func init() {
	banks := make(map[string]*BankPro)
	banks["cn_psbc_dc_19"] = &BankPro{
		Bank{
			"psbc","cn","中国邮政储蓄银行","","","",
			[]int{621096,621098,622150,622151,622181,622188,622199,955100,621095,620062,621285,621798,621799,621797,620529,621622,621599,621674,623218,623219},
		},
	"dc", 19,
	}
	banks["cn_psbc_dc_19"] = &BankPro{
		Bank{
			"psbc","cn","中国邮政储蓄银行","","","",
			[]int{62215049,62215050,62215051,62218850,62218851,62218849},
		},
	"dc", 19,
	}
	banks["cn_psbc_cc_16"] = &BankPro{
		Bank{
			"psbc","cn","中国邮政储蓄银行","","","",
			[]int{622812,622810,622811,628310,625919},
		},
	"cc", 16,
	}
	banks["cn_icbc_dc_18"] = &BankPro{
		Bank{
			"icbc","cn","中国工商银行","Industrial and Commercial Bank of China","http://www.icbc.com.cn","#c90000",
			[]int{620200,620302,620402,620403,620404,620406,620407,620409,620410,620411,620412,620502,620503,620405,620408,620512,620602,620604,620607,620611,620612,620704,620706,620707,620708,620709,620710,620609,620712,620713,620714,620802,620711,620904,620905,621001,620902,621103,621105,621106,621107,621102,621203,621204,621205,621206,621207,621208,621209,621210,621302,621303,621202,621305,621306,621307,621309,621311,621313,621211,621315,621304,621402,621404,621405,621406,621407,621408,621409,621410,621502,621317,621511,621602,621603,621604,621605,621608,621609,621610,621611,621612,621613,621614,621615,621616,621617,621607,621606,621804,621807,621813,621814,621817,621901,621904,621905,621906,621907,621908,621909,621910,621911,621912,621913,621915,622002,621903,622004,622005,622006,622007,622008,622010,622011,622012,621914,622015,622016,622003,622018,622019,622020,622102,622103,622104,622105,622013,622111,622114,622017,622110,622303,622304,622305,622306,622307,622308,622309,622314,622315,622317,622302,622402,622403,622404,622313,622504,622505,622509,622513,622517,622502,622604,622605,622606,622510,622703,622715,622806,622902,622903,622706,623002,623006,623008,623011,623012,622904,623015,623100,623202,623301,623400,623500,623602,623803,623901,623014,624100,624200,624301,624402,623700,624000},
		},
	"dc", 18,
	}
	banks["cn_icbc_dc_19"] = &BankPro{
		Bank{
			"icbc","cn","中国工商银行","Industrial and Commercial Bank of China","http://www.icbc.com.cn","#c90000",
			[]int{622200,622202,622203,622208,621225,620058,621281,900000,621558,621559,621722,621723,620086,621226,621618,620516,621227,621288,621721,900010,623062,621670,621720,621379,621240,621724,621762,621414,621375,622926,622927,622928,622929,622930,622931,621733,621732,621372,621369,621763},
		},
	"dc", 19,
	}
	banks["cn_icbc_dc_16"] = &BankPro{
		Bank{
			"icbc","cn","中国工商银行","Industrial and Commercial Bank of China","http://www.icbc.com.cn","#c90000",
			[]int{402791,427028,427038,548259,621376,621423,621428,621434,621761,621749,621300,621378,622944,622949,621371,621730,621734,621433,621370,621764,621464,621765,621750,621377,621367,621374,621731,621781},
		},
	"dc", 16,
	}
	banks["cn_icbc_dc_19"] = &BankPro{
		Bank{
			"icbc","cn","中国工商银行","Industrial and Commercial Bank of China","http://www.icbc.com.cn","#c90000",
			[]int{9558},
		},
	"dc", 19,
	}
	banks["cn_icbc_cc_15"] = &BankPro{
		Bank{
			"icbc","cn","中国工商银行","Industrial and Commercial Bank of China","http://www.icbc.com.cn","#c90000",
			[]int{370246,370248,370249,370247,370267,374738,374739},
		},
	"cc", 15,
	}
	banks["cn_icbc_cc_16"] = &BankPro{
		Bank{
			"icbc","cn","中国工商银行","Industrial and Commercial Bank of China","http://www.icbc.com.cn","#c90000",
			[]int{427010,427018,427019,427020,427029,427030,427039,438125,438126,451804,451810,451811,458071,489734,489735,489736,510529,427062,524091,427064,530970,530990,558360,524047,525498,622230,622231,622232,622233,622234,622235,622237,622239,622240,622245,622238,451804,451810,451811,458071,628288,628286,622206,526836,513685,543098,458441,622246,544210,548943,356879,356880,356881,356882,528856,625330,625331,625332,622236,524374,550213,625929,625927,625939,625987,625930,625114,622159,625021,625022,625932,622889,625900,625915,625916,622171,625931,625113,625928,625914,625986,625925,625921,625926,625942,622158,625917,625922,625934,625933,625920,625924,625017,625018,625019},
		},
	"cc", 16,
	}
	banks["cn_icbc_cc_16"] = &BankPro{
		Bank{
			"icbc","cn","中国工商银行","Industrial and Commercial Bank of China","http://www.icbc.com.cn","#c90000",
			[]int{45806,53098,45806,53098},
		},
	"cc", 16,
	}
	banks["cn_icbc_scc_16"] = &BankPro{
		Bank{
			"icbc","cn","中国工商银行","Industrial and Commercial Bank of China","http://www.icbc.com.cn","#c90000",
			[]int{622210,622211,622212,622213,622214,622220,622223,622225,622229,622215,622224},
		},
	"scc", 16,
	}
	banks["cn_icbc_pc_16"] = &BankPro{
		Bank{
			"icbc","cn","中国工商银行","Industrial and Commercial Bank of China","http://www.icbc.com.cn","#c90000",
			[]int{620054,620142,620184,620030,620050,620143,620149,620124,620183,620094,620186,620148,620185},
		},
	"pc", 16,
	}
	banks["cn_icbc_pc_19"] = &BankPro{
		Bank{
			"icbc","cn","中国工商银行","Industrial and Commercial Bank of China","http://www.icbc.com.cn","#c90000",
			[]int{620114,620187,620046},
		},
	"pc", 19,
	}
	banks["cn_abc_dc_19"] = &BankPro{
		Bank{
			"abc","cn","中国农业银行","Agricultural Bank of China","http://www.abchina.com/","#319c8b",
			[]int{622841,622824,622826,622848,620059,621282,622828,622823,621336,621619,622821,622822,622825,622827,622845,622849,623018,623206,621671,622840,622843,622844,622846,622847,620501},
		},
	"dc", 19,
	}
	banks["cn_abc_dc_19"] = &BankPro{
		Bank{
			"abc","cn","中国农业银行","Agricultural Bank of China","http://www.abchina.com/","#319c8b",
			[]int{95595,95596,95597,95598,95599},
		},
	"dc", 19,
	}
	banks["cn_abc_dc_19"] = &BankPro{
		Bank{
			"abc","cn","中国农业银行","Agricultural Bank of China","http://www.abchina.com/","#319c8b",
			[]int{103},
		},
	"dc", 19,
	}
	banks["cn_abc_cc_16"] = &BankPro{
		Bank{
			"abc","cn","中国农业银行","Agricultural Bank of China","http://www.abchina.com/","#319c8b",
			[]int{403361,404117,404118,404119,404120,404121,463758,519412,519413,520082,520083,552599,558730,514027,622836,622837,628268,625996,625998,625997,622838,625336,625826,625827,544243,548478,628269},
		},
	"cc", 16,
	}
	banks["cn_abc_scc_16"] = &BankPro{
		Bank{
			"abc","cn","中国农业银行","Agricultural Bank of China","http://www.abchina.com/","#319c8b",
			[]int{622820,622830},
		},
	"scc", 16,
	}
	banks["cn_boc_dc_19"] = &BankPro{
		Bank{
			"boc","cn","中国银行","Bank of China","http://www.boc.cn/","#a71e32",
			[]int{621660,621661,621662,621663,621665,621667,621668,621669,621666,456351,601382,621256,621212,621283,620061,621725,621330,621331,621332,621333,621297,621568,621569,621672,623208,621620,621756,621757,621758,621759,621785,621786,621787,621788,621789,621790,622273,622274,622771,622772,622770,621741,621041},
		},
	"dc", 19,
	}
	banks["cn_boc_dc_16"] = &BankPro{
		Bank{
			"boc","cn","中国银行","Bank of China","http://www.boc.cn/","#a71e32",
			[]int{621293,621294,621342,621343,621364,621394,621648,621248,621215,621249,621231,621638,621334,621395,623040,622348},
		},
	"dc", 16,
	}
	banks["cn_boc_cc_16"] = &BankPro{
		Bank{
			"boc","cn","中国银行","Bank of China","http://www.boc.cn/","#a71e32",
			[]int{625908,625910,625909,356833,356835,409665,409666,409668,409669,409670,409671,409672,512315,512316,512411,512412,514957,409667,438088,552742,553131,514958,622760,628388,518377,622788,628313,628312,622750,622751,625145,622479,622480,622789,625140,622346,622347},
		},
	"cc", 16,
	}
	banks["cn_boc_scc_16"] = &BankPro{
		Bank{
			"boc","cn","中国银行","Bank of China","http://www.boc.cn/","#a71e32",
			[]int{518378,518379,518474,518475,518476,524865,525745,525746,547766,558868,622752,622753,622755,524864,622757,622758,622759,622761,622762,622763,622756,622754,622764,622765,558869,625905,625906,625907,625333},
		},
	"scc", 16,
	}
	banks["cn_boc_scc_16"] = &BankPro{
		Bank{
			"boc","cn","中国银行","Bank of China","http://www.boc.cn/","#a71e32",
			[]int{53591,49102,377677},
		},
	"scc", 16,
	}
	banks["cn_boc_pc_16"] = &BankPro{
		Bank{
			"boc","cn","中国银行","Bank of China","http://www.boc.cn/","#a71e32",
			[]int{620514,620025,620026,620210,620211,620019,620035,620202,620203,620048,620515,920000},
		},
	"pc", 16,
	}
	banks["cn_boc_pc_19"] = &BankPro{
		Bank{
			"boc","cn","中国银行","Bank of China","http://www.boc.cn/","#a71e32",
			[]int{620040,620531,620513,921000,620038},
		},
	"pc", 19,
	}
	banks["cn_ccb_dc_19"] = &BankPro{
		Bank{
			"ccb","cn","中国建设银行","China Construction Bank","http://www.ccb.com/","#0c69b0",
			[]int{621284,436742,589970,620060,621081,621467,621598,621621,621700,622280,622700,623211,623668},
		},
	"dc", 19,
	}
	banks["cn_ccb_dc_16"] = &BankPro{
		Bank{
			"ccb","cn","中国建设银行","China Construction Bank","http://www.ccb.com/","#0c69b0",
			[]int{421349,434061,434062,524094,526410,552245,621080,621082,621466,621488,621499,622966,622988,622382,621487,621083,621084,620107},
		},
	"dc", 16,
	}
	banks["cn_ccb_dc_19"] = &BankPro{
		Bank{
			"ccb","cn","中国建设银行","China Construction Bank","http://www.ccb.com/","#0c69b0",
			[]int{436742193,622280193},
		},
	"dc", 19,
	}
	banks["cn_ccb_cc_18"] = &BankPro{
		Bank{
			"ccb","cn","中国建设银行","China Construction Bank","http://www.ccb.com/","#0c69b0",
			[]int{553242},
		},
	"cc", 18,
	}
	banks["cn_ccb_cc_16"] = &BankPro{
		Bank{
			"ccb","cn","中国建设银行","China Construction Bank","http://www.ccb.com/","#0c69b0",
			[]int{625362,625363,628316,628317,356896,356899,356895,436718,436738,436745,436748,489592,531693,532450,532458,544887,552801,557080,558895,559051,622166,622168,622708,625964,625965,625966,628266,628366,622381,622675,622676,622677},
		},
	"cc", 16,
	}
	banks["cn_ccb_cc_18"] = &BankPro{
		Bank{
			"ccb","cn","中国建设银行","China Construction Bank","http://www.ccb.com/","#0c69b0",
			[]int{5453242,5491031,5544033},
		},
	"cc", 18,
	}
	banks["cn_ccb_scc_16"] = &BankPro{
		Bank{
			"ccb","cn","中国建设银行","China Construction Bank","http://www.ccb.com/","#0c69b0",
			[]int{622725,622728,436728,453242,491031,544033,622707,625955,625956},
		},
	"scc", 16,
	}
	banks["cn_ccb_scc_16"] = &BankPro{
		Bank{
			"ccb","cn","中国建设银行","China Construction Bank","http://www.ccb.com/","#0c69b0",
			[]int{53242,53243},
		},
	"scc", 16,
	}
	banks["cn_comm_dc_19"] = &BankPro{
		Bank{
			"comm","cn","中国交通银行","","","",
			[]int{622261,622260,622262,621002,621069,621436,621335},
		},
	"dc", 19,
	}
	banks["cn_comm_dc_16"] = &BankPro{
		Bank{
			"comm","cn","中国交通银行","","","",
			[]int{620013},
		},
	"dc", 16,
	}
	banks["cn_comm_dc_17"] = &BankPro{
		Bank{
			"comm","cn","中国交通银行","","","",
			[]int{405512,601428,405512,601428,622258,622259,405512,601428},
		},
	"dc", 17,
	}
	banks["cn_comm_cc_16"] = &BankPro{
		Bank{
			"comm","cn","中国交通银行","","","",
			[]int{49104,53783},
		},
	"cc", 16,
	}
	banks["cn_comm_cc_16"] = &BankPro{
		Bank{
			"comm","cn","中国交通银行","","","",
			[]int{434910,458123,458124,520169,522964,552853,622250,622251,521899,622253,622656,628216,622252,955590,955591,955592,955593,628218,625028,625029},
		},
	"cc", 16,
	}
	banks["cn_comm_scc_16"] = &BankPro{
		Bank{
			"comm","cn","中国交通银行","","","",
			[]int{622254,622255,622256,622257,622284},
		},
	"scc", 16,
	}
	banks["cn_comm_pc_19"] = &BankPro{
		Bank{
			"comm","cn","中国交通银行","","","",
			[]int{620021,620521},
		},
	"pc", 19,
	}
	banks["cn_cmb_dc_17"] = &BankPro{
		Bank{
			"cmb","cn","招商银行","China Merchants Bank","http://www.cmbchina.com","#b0110c",
			[]int{402658,410062,468203,512425,524011,622580,622588,622598,622609,95555,621286,621483,621485,621486,621299,10},
		},
	"dc", 17,
	}
	banks["cn_cmb_dc_15"] = &BankPro{
		Bank{
			"cmb","cn","招商银行","China Merchants Bank","http://www.cmbchina.com","#b0110c",
			[]int{690755},
		},
	"dc", 15,
	}
	banks["cn_cmb_dc_18"] = &BankPro{
		Bank{
			"cmb","cn","招商银行","China Merchants Bank","http://www.cmbchina.com","#b0110c",
			[]int{690755},
		},
	"dc", 18,
	}
	banks["cn_cmb_cc_16"] = &BankPro{
		Bank{
			"cmb","cn","招商银行","China Merchants Bank","http://www.cmbchina.com","#b0110c",
			[]int{356885,356886,356887,356888,356890,439188,439227,479228,479229,521302,356889,545620,545621,545947,545948,552534,552587,622575,622576,622577,622578,622579,545619,622581,622582,545623,628290,439225,518710,518718,628362,439226,628262,625802,625803},
		},
	"cc", 16,
	}
	banks["cn_cmb_cc_15"] = &BankPro{
		Bank{
			"cmb","cn","招商银行","China Merchants Bank","http://www.cmbchina.com","#b0110c",
			[]int{370285,370286,370287,370289},
		},
	"cc", 15,
	}
	banks["cn_cmb_pc_19"] = &BankPro{
		Bank{
			"cmb","cn","招商银行","China Merchants Bank","http://www.cmbchina.com","#b0110c",
			[]int{620520},
		},
	"pc", 19,
	}
	banks["cn_cmbc_dc_16"] = &BankPro{
		Bank{
			"cmbc","cn","中国民生银行","","","",
			[]int{622615,622616,622618,622622,622617,622619,415599,421393,421865,427570,427571,472067,472068,622620},
		},
	"dc", 16,
	}
	banks["cn_cmbc_cc_16"] = &BankPro{
		Bank{
			"cmbc","cn","中国民生银行","","","",
			[]int{545392,545393,545431,545447,356859,356857,407405,421869,421870,421871,512466,356856,528948,552288,622600,622601,622602,517636,622621,628258,556610,622603,464580,464581,523952,545217,553161,356858,622623,625912,625913,625911},
		},
	"cc", 16,
	}
	banks["cn_cmbc_cc_15"] = &BankPro{
		Bank{
			"cmbc","cn","中国民生银行","","","",
			[]int{377155,377152,377153,377158},
		},
	"cc", 15,
	}
	banks["cn_ceb_dc_16"] = &BankPro{
		Bank{
			"ceb","cn","中国光大银行","China Everbright Bank","http://www.cebbank.cn/","#691585",
			[]int{303},
		},
	"dc", 16,
	}
	banks["cn_ceb_dc_16"] = &BankPro{
		Bank{
			"ceb","cn","中国光大银行","China Everbright Bank","http://www.cebbank.cn/","#691585",
			[]int{90030},
		},
	"dc", 16,
	}
	banks["cn_ceb_dc_19"] = &BankPro{
		Bank{
			"ceb","cn","中国光大银行","China Everbright Bank","http://www.cebbank.cn/","#691585",
			[]int{620535},
		},
	"dc", 19,
	}
	banks["cn_ceb_dc_16"] = &BankPro{
		Bank{
			"ceb","cn","中国光大银行","China Everbright Bank","http://www.cebbank.cn/","#691585",
			[]int{620085,622660,622662,622663,622664,622665,622666,622667,622669,622670,622671,622672,622668,622661,622674,622673,620518,621489,621492},
		},
	"dc", 16,
	}
	banks["cn_ceb_cc_16"] = &BankPro{
		Bank{
			"ceb","cn","中国光大银行","China Everbright Bank","http://www.cebbank.cn/","#691585",
			[]int{356837,356838,486497,622657,622685,622659,622687,625978,625980,625981,625979,356839,356840,406252,406254,425862,481699,524090,543159,622161,622570,622650,622655,622658,625975,625977,628201,628202,625339,625976},
		},
	"cc", 16,
	}
	banks["cn_citic_dc_16"] = &BankPro{
		Bank{
			"citic","cn","中信银行","China Citic Bank","http://bank.ecitic.com/","#c8151c",
			[]int{433670,433680,442729,442730,620082,622690,622691,622692,622696,622698,622998,622999,433671,968807,968808,968809,621771,621767,621768,621770,621772,621773,622453,622456},
		},
	"dc", 16,
	}
	banks["cn_citic_dc_17"] = &BankPro{
		Bank{
			"citic","cn","中信银行","China Citic Bank","http://bank.ecitic.com/","#c8151c",
			[]int{622459},
		},
	"dc", 17,
	}
	banks["cn_citic_cc_15"] = &BankPro{
		Bank{
			"citic","cn","中信银行","China Citic Bank","http://bank.ecitic.com/","#c8151c",
			[]int{376968,376969,376966},
		},
	"cc", 15,
	}
	banks["cn_citic_cc_16"] = &BankPro{
		Bank{
			"citic","cn","中信银行","China Citic Bank","http://bank.ecitic.com/","#c8151c",
			[]int{400360,403391,403392,404158,404159,404171,404172,404173,404174,404157,433667,433668,433669,514906,403393,520108,433666,558916,622678,622679,622680,622688,622689,628206,556617,628209,518212,628208,356390,356391,356392,622916,622918,622919},
		},
	"cc", 16,
	}
	banks["cn_hxbank_dc_16"] = &BankPro{
		Bank{
			"hxbank","cn","华夏银行","","","",
			[]int{622630,622631,622632,622633,999999,621222,623020,623021,623022,623023},
		},
	"dc", 16,
	}
	banks["cn_hxbank_cc_16"] = &BankPro{
		Bank{
			"hxbank","cn","华夏银行","","","",
			[]int{523959,528709,539867,539868,622637,622638,628318,528708,622636,625967,625968,625969},
		},
	"cc", 16,
	}
	banks["cn_spabank_dc_19"] = &BankPro{
		Bank{
			"spabank","cn","深发/平安银行","","","",
			[]int{621626,623058},
		},
	"dc", 19,
	}
	banks["cn_spabank_dc_16"] = &BankPro{
		Bank{
			"spabank","cn","深发/平安银行","","","",
			[]int{602907,622986,622989,622298,627069,627068,627066,627067,412963,415752,415753,622535,622536,622538,622539,998800,412962,622983},
		},
	"dc", 16,
	}
	banks["cn_spabank_cc_16"] = &BankPro{
		Bank{
			"spabank","cn","深发/平安银行","","","",
			[]int{531659,622157,528020,622155,622156,526855,356869,356868,625360,625361,628296,435744,435745,483536,622525,622526,998801,998802},
		},
	"cc", 16,
	}
	banks["cn_spabank_pc_16"] = &BankPro{
		Bank{
			"spabank","cn","深发/平安银行","","","",
			[]int{620010},
		},
	"pc", 16,
	}
	banks["cn_cib_dc_18"] = &BankPro{
		Bank{
			"cib","cn","兴业银行","Industrial Bank Co. Ltd.","http://www.cib.com.cn/","#004097",
			[]int{438589},
		},
	"dc", 18,
	}
	banks["cn_cib_dc_16"] = &BankPro{
		Bank{
			"cib","cn","兴业银行","Industrial Bank Co. Ltd.","http://www.cib.com.cn/","#004097",
			[]int{90592},
		},
	"dc", 16,
	}
	banks["cn_cib_dc_18"] = &BankPro{
		Bank{
			"cib","cn","兴业银行","Industrial Bank Co. Ltd.","http://www.cib.com.cn/","#004097",
			[]int{966666,622909,438588,622908},
		},
	"dc", 18,
	}
	banks["cn_cib_cc_16"] = &BankPro{
		Bank{
			"cib","cn","兴业银行","Industrial Bank Co. Ltd.","http://www.cib.com.cn/","#004097",
			[]int{461982,486493,486494,486861,523036,451289,527414,528057,622901,622902,622922,628212,451290,524070,625084,625085,625086,625087,548738,549633,552398,625082,625083,625960,625961,625962,625963},
		},
	"cc", 16,
	}
	banks["cn_cib_pc_16"] = &BankPro{
		Bank{
			"cib","cn","兴业银行","Industrial Bank Co. Ltd.","http://www.cib.com.cn/","#004097",
			[]int{620010},
		},
	"pc", 16,
	}
	banks["cn_shbank_dc_18"] = &BankPro{
		Bank{
			"shbank","cn","上海银行","","","",
			[]int{621050,622172,622985,622987,620522,622267,622278,622279,622468,622892,940021},
		},
	"dc", 18,
	}
	banks["cn_shbank_dc_16"] = &BankPro{
		Bank{
			"shbank","cn","上海银行","","","",
			[]int{438600},
		},
	"dc", 16,
	}
	banks["cn_shbank_cc_16"] = &BankPro{
		Bank{
			"shbank","cn","上海银行","","","",
			[]int{356827,356828,356830,402673,402674,486466,519498,520131,524031,548838,622148,622149,622268,356829,622300,628230,622269,625099,625953},
		},
	"cc", 16,
	}
	banks["cn_spdb_dc_16"] = &BankPro{
		Bank{
			"spdb","cn","浦东发展银行","Shanghai Pudong Development Bank","http://www.spdb.com.cn","#072154",
			[]int{622516,622517,622518,622521,622522,622523,984301,984303,621352,621793,621795,621796,621351,621390,621792,621791},
		},
	"dc", 16,
	}
	banks["cn_spdb_dc_16"] = &BankPro{
		Bank{
			"spdb","cn","浦东发展银行","Shanghai Pudong Development Bank","http://www.spdb.com.cn","#072154",
			[]int{84301,84336,84373,84385,84390,87000,87010,87030,87040,84380,84361,87050,84342},
		},
	"dc", 16,
	}
	banks["cn_spdb_cc_16"] = &BankPro{
		Bank{
			"spdb","cn","浦东发展银行","Shanghai Pudong Development Bank","http://www.spdb.com.cn","#072154",
			[]int{356851,356852,404738,404739,456418,498451,515672,356850,517650,525998,622177,622277,628222,622500,628221,622176,622276,622228,625957,625958,625993,625831},
		},
	"cc", 16,
	}
	banks["cn_spdb_scc_16"] = &BankPro{
		Bank{
			"spdb","cn","浦东发展银行","Shanghai Pudong Development Bank","http://www.spdb.com.cn","#072154",
			[]int{622520,622519},
		},
	"scc", 16,
	}
	banks["cn_spdb_pc_19"] = &BankPro{
		Bank{
			"spdb","cn","浦东发展银行","Shanghai Pudong Development Bank","http://www.spdb.com.cn","#072154",
			[]int{620530},
		},
	"pc", 19,
	}
	banks["cn_gdb_dc_16"] = &BankPro{
		Bank{
			"gdb","cn","广发银行","","","",
			[]int{622516,622517,622518,622521,622522,622523,984301,984303,621352,621793,621795,621796,621351,621390,621792,621791},
		},
	"dc", 16,
	}
	banks["cn_gdb_dc_19"] = &BankPro{
		Bank{
			"gdb","cn","广发银行","","","",
			[]int{622568,6858001,6858009,621462},
		},
	"dc", 19,
	}
	banks["cn_gdb_dc_19"] = &BankPro{
		Bank{
			"gdb","cn","广发银行","","","",
			[]int{9111},
		},
	"dc", 19,
	}
	banks["cn_gdb_cc_16"] = &BankPro{
		Bank{
			"gdb","cn","广发银行","","","",
			[]int{406365,406366,428911,436768,436769,436770,487013,491032,491033,491034,491035,491036,491037,491038,436771,518364,520152,520382,541709,541710,548844,552794,493427,622555,622556,622557,622558,622559,622560,528931,558894,625072,625071,628260,628259,625805,625806,625807,625808,625809,625810},
		},
	"cc", 16,
	}
	banks["cn_gdb_cc_19"] = &BankPro{
		Bank{
			"gdb","cn","广发银行","","","",
			[]int{685800,6858000},
		},
	"cc", 19,
	}
	banks["cn_bohaib_dc_16"] = &BankPro{
		Bank{
			"bohaib","cn","渤海银行","","","",
			[]int{621268,622684,622884,621453},
		},
	"dc", 16,
	}
	banks["cn_gcb_dc_19"] = &BankPro{
		Bank{
			"gcb","cn","广州银行","","","",
			[]int{603445,622467,940016,621463},
		},
	"dc", 19,
	}
	banks["cn_jhbank_dc_16"] = &BankPro{
		Bank{
			"jhbank","cn","金华银行","","","",
			[]int{622449,940051},
		},
	"dc", 16,
	}
	banks["cn_jhbank_cc_16"] = &BankPro{
		Bank{
			"jhbank","cn","金华银行","","","",
			[]int{622450,628204},
		},
	"cc", 16,
	}
	banks["cn_wzcb_dc_16"] = &BankPro{
		Bank{
			"wzcb","cn","温州银行","","","",
			[]int{621977},
		},
	"dc", 16,
	}
	banks["cn_wzcb_cc_16"] = &BankPro{
		Bank{
			"wzcb","cn","温州银行","","","",
			[]int{622868,622899,628255},
		},
	"cc", 16,
	}
	banks["cn_hsbank_dc_19"] = &BankPro{
		Bank{
			"hsbank","cn","徽商银行","","","",
			[]int{622877,622879,621775,623203},
		},
	"dc", 19,
	}
	banks["cn_hsbank_dc_17"] = &BankPro{
		Bank{
			"hsbank","cn","徽商银行","","","",
			[]int{603601,622137,622327,622340,622366},
		},
	"dc", 17,
	}
	banks["cn_hsbank_cc_16"] = &BankPro{
		Bank{
			"hsbank","cn","徽商银行","","","",
			[]int{628251,622651,625828},
		},
	"cc", 16,
	}
	banks["cn_jsbank_dc_19"] = &BankPro{
		Bank{
			"jsbank","cn","江苏银行","","","",
			[]int{621076,622173,622131,621579,622876},
		},
	"dc", 19,
	}
	banks["cn_jsbank_dc_16"] = &BankPro{
		Bank{
			"jsbank","cn","江苏银行","","","",
			[]int{504923,622422,622447,940076},
		},
	"dc", 16,
	}
	banks["cn_jsbank_cc_16"] = &BankPro{
		Bank{
			"jsbank","cn","江苏银行","","","",
			[]int{628210,622283,625902},
		},
	"cc", 16,
	}
	banks["cn_njcb_dc_16"] = &BankPro{
		Bank{
			"njcb","cn","南京银行","","","",
			[]int{621777,622305,621259},
		},
	"dc", 16,
	}
	banks["cn_njcb_cc_16"] = &BankPro{
		Bank{
			"njcb","cn","南京银行","","","",
			[]int{622303,628242,622595,622596},
		},
	"cc", 16,
	}
	banks["cn_nbbank_dc_16"] = &BankPro{
		Bank{
			"nbbank","cn","宁波银行","","","",
			[]int{621279,622281,622316,940022},
		},
	"dc", 16,
	}
	banks["cn_nbbank_dc_19"] = &BankPro{
		Bank{
			"nbbank","cn","宁波银行","","","",
			[]int{621418},
		},
	"dc", 19,
	}
	banks["cn_nbbank_cc_16"] = &BankPro{
		Bank{
			"nbbank","cn","宁波银行","","","",
			[]int{625903,622778,628207,512431,520194,622282,622318},
		},
	"cc", 16,
	}
	banks["cn_bjbank_dc_16"] = &BankPro{
		Bank{
			"bjbank","cn","北京银行","","","",
			[]int{623111,421317,422161,602969,422160,621030,621420,621468},
		},
	"dc", 16,
	}
	banks["cn_bjbank_cc_16"] = &BankPro{
		Bank{
			"bjbank","cn","北京银行","","","",
			[]int{522001,622163,622853,628203,622851,622852},
		},
	"cc", 16,
	}
	banks["cn_bjrcb_dc_19"] = &BankPro{
		Bank{
			"bjrcb","cn","北京农村商业银行","","","",
			[]int{620088,621068,622138,621066,621560},
		},
	"dc", 19,
	}
	banks["cn_bjrcb_cc_16"] = &BankPro{
		Bank{
			"bjrcb","cn","北京农村商业银行","","","",
			[]int{625526,625186,628336},
		},
	"cc", 16,
	}
	banks["cn_hsbc_dc_16"] = &BankPro{
		Bank{
			"hsbc","cn","汇丰银行","","","",
			[]int{622946},
		},
	"dc", 16,
	}
	banks["cn_hsbc_dc_17"] = &BankPro{
		Bank{
			"hsbc","cn","汇丰银行","","","",
			[]int{622406,621442},
		},
	"dc", 17,
	}
	banks["cn_hsbc_dc_19"] = &BankPro{
		Bank{
			"hsbc","cn","汇丰银行","","","",
			[]int{622407,621443},
		},
	"dc", 19,
	}
	banks["cn_hsbc_cc_16"] = &BankPro{
		Bank{
			"hsbc","cn","汇丰银行","","","",
			[]int{622360,622361,625034,625096,625098},
		},
	"cc", 16,
	}
	banks["cn_scb_dc_16"] = &BankPro{
		Bank{
			"scb","cn","渣打银行","","","",
			[]int{622948,621740,622942,622994},
		},
	"dc", 16,
	}
	banks["cn_scb_cc_16"] = &BankPro{
		Bank{
			"scb","cn","渣打银行","","","",
			[]int{622482,622483,622484},
		},
	"cc", 16,
	}
	banks["cn_citi_dc_16"] = &BankPro{
		Bank{
			"citi","cn","花旗银行","","","",
			[]int{621062,621063},
		},
	"dc", 16,
	}
	banks["cn_citi_cc_16"] = &BankPro{
		Bank{
			"citi","cn","花旗银行","","","",
			[]int{625076,625077,625074,625075,622371,625091},
		},
	"cc", 16,
	}
	banks["cn_hkbea_dc_19"] = &BankPro{
		Bank{
			"hkbea","cn","东亚银行","","","",
			[]int{622933,622938,623031,622943,621411},
		},
	"dc", 19,
	}
	banks["cn_hkbea_cc_16"] = &BankPro{
		Bank{
			"hkbea","cn","东亚银行","","","",
			[]int{622372,622471,622472,622265,622266,625972,625973},
		},
	"cc", 16,
	}
	banks["cn_hkbea_cc_17"] = &BankPro{
		Bank{
			"hkbea","cn","东亚银行","","","",
			[]int{622365},
		},
	"cc", 17,
	}
	banks["cn_ghb_dc_19"] = &BankPro{
		Bank{
			"ghb","cn","广东华兴银行","","","",
			[]int{621469,621625},
		},
	"dc", 19,
	}
	banks["cn_srcb_dc_16"] = &BankPro{
		Bank{
			"srcb","cn","深圳农村商业银行","","","",
			[]int{622128,622129,623035},
		},
	"dc", 16,
	}
	banks["cn_gzrcu_dc_18"] = &BankPro{
		Bank{
			"gzrcu","cn","广州农村商业银行股份有限公司","","","",
			[]int{909810,940035,621522,622439},
		},
	"dc", 18,
	}
	banks["cn_drcbcl_dc_19"] = &BankPro{
		Bank{
			"drcbcl","cn","东莞农村商业银行","","","",
			[]int{622328,940062,623038},
		},
	"dc", 19,
	}
	banks["cn_drcbcl_cc_16"] = &BankPro{
		Bank{
			"drcbcl","cn","东莞农村商业银行","","","",
			[]int{625288,625888},
		},
	"cc", 16,
	}
	banks["cn_bod_dc_16"] = &BankPro{
		Bank{
			"bod","cn","东莞市商业银行","","","",
			[]int{622333,940050},
		},
	"dc", 16,
	}
	banks["cn_bod_dc_19"] = &BankPro{
		Bank{
			"bod","cn","东莞市商业银行","","","",
			[]int{621439,623010},
		},
	"dc", 19,
	}
	banks["cn_bod_cc_16"] = &BankPro{
		Bank{
			"bod","cn","东莞市商业银行","","","",
			[]int{622888},
		},
	"cc", 16,
	}
	banks["cn_gdrcc_dc_16"] = &BankPro{
		Bank{
			"gdrcc","cn","广东省农村信用社联合社","","","",
			[]int{622302},
		},
	"dc", 16,
	}
	banks["cn_gdrcc_dc_19"] = &BankPro{
		Bank{
			"gdrcc","cn","广东省农村信用社联合社","","","",
			[]int{622477,622509,622510,622362,621018,621518},
		},
	"dc", 19,
	}
	banks["cn_dsb_dc_16"] = &BankPro{
		Bank{
			"dsb","cn","大新银行","","","",
			[]int{622297,621277},
		},
	"dc", 16,
	}
	banks["cn_dsb_dc_17"] = &BankPro{
		Bank{
			"dsb","cn","大新银行","","","",
			[]int{622375,622489},
		},
	"dc", 17,
	}
	banks["cn_dsb_cc_16"] = &BankPro{
		Bank{
			"dsb","cn","大新银行","","","",
			[]int{622293,622295,622296,622373,622451,622294,625940},
		},
	"cc", 16,
	}
	banks["cn_whb_dc_16"] = &BankPro{
		Bank{
			"whb","cn","永亨银行","","","",
			[]int{622871,622958,622963,622957,622861,622932,622862,621298},
		},
	"dc", 16,
	}
	banks["cn_whb_cc_16"] = &BankPro{
		Bank{
			"whb","cn","永亨银行","","","",
			[]int{622798,625010,622775,622785},
		},
	"cc", 16,
	}
	banks["cn_dbs_dc_19"] = &BankPro{
		Bank{
			"dbs","cn","星展银行香港有限公司","","","",
			[]int{621016,621015},
		},
	"dc", 19,
	}
	banks["cn_dbs_dc_16"] = &BankPro{
		Bank{
			"dbs","cn","星展银行香港有限公司","","","",
			[]int{622487,622490,622491,622492},
		},
	"dc", 16,
	}
	banks["cn_dbs_dc_17"] = &BankPro{
		Bank{
			"dbs","cn","星展银行香港有限公司","","","",
			[]int{622487,622490,622491,622492,621744,621745,621746,621747},
		},
	"dc", 17,
	}
	banks["cn_egbank_dc_19"] = &BankPro{
		Bank{
			"egbank","cn","恒丰银行","","","",
			[]int{623078},
		},
	"dc", 19,
	}
	banks["cn_egbank_dc_17"] = &BankPro{
		Bank{
			"egbank","cn","恒丰银行","","","",
			[]int{622384,940034},
		},
	"dc", 17,
	}
	banks["cn_tccb_dc_18"] = &BankPro{
		Bank{
			"tccb","cn","天津市商业银行","","","",
			[]int{940015,622331},
		},
	"dc", 18,
	}
	banks["cn_tccb_dc_18"] = &BankPro{
		Bank{
			"tccb","cn","天津市商业银行","","","",
			[]int{6091201},
		},
	"dc", 18,
	}
	banks["cn_tccb_cc_16"] = &BankPro{
		Bank{
			"tccb","cn","天津市商业银行","","","",
			[]int{622426,628205},
		},
	"cc", 16,
	}
	banks["cn_czbank_dc_19"] = &BankPro{
		Bank{
			"czbank","cn","浙商银行","","","",
			[]int{621019,622309,621019},
		},
	"dc", 19,
	}
	banks["cn_czbank_dc_19"] = &BankPro{
		Bank{
			"czbank","cn","浙商银行","","","",
			[]int{6223091100,6223092900,6223093310,6223093320,6223093330,6223093370,6223093380,6223096510,6223097910},
		},
	"dc", 19,
	}
	banks["cn_ncb_dc_19"] = &BankPro{
		Bank{
			"ncb","cn","南洋商业银行","","","",
			[]int{621213,621289,621290,621291,621292,621042,621743},
		},
	"dc", 19,
	}
	banks["cn_ncb_dc_16"] = &BankPro{
		Bank{
			"ncb","cn","南洋商业银行","","","",
			[]int{623041,622351},
		},
	"dc", 16,
	}
	banks["cn_ncb_cc_16"] = &BankPro{
		Bank{
			"ncb","cn","南洋商业银行","","","",
			[]int{625046,625044,625058,622349,622350},
		},
	"cc", 16,
	}
	banks["cn_ncb_pc_16"] = &BankPro{
		Bank{
			"ncb","cn","南洋商业银行","","","",
			[]int{620208,620209,625093,625095},
		},
	"pc", 16,
	}
	banks["cn_xmbank_dc_16"] = &BankPro{
		Bank{
			"xmbank","cn","厦门银行","","","",
			[]int{622393,940023},
		},
	"dc", 16,
	}
	banks["cn_xmbank_dc_18"] = &BankPro{
		Bank{
			"xmbank","cn","厦门银行","","","",
			[]int{6886592},
		},
	"dc", 18,
	}
	banks["cn_xmbank_dc_19"] = &BankPro{
		Bank{
			"xmbank","cn","厦门银行","","","",
			[]int{623019,621600},
		},
	"dc", 19,
	}
	banks["cn_fjhxbc_dc_16"] = &BankPro{
		Bank{
			"fjhxbc","cn","福建海峡银行","","","",
			[]int{622388},
		},
	"dc", 16,
	}
	banks["cn_fjhxbc_dc_18"] = &BankPro{
		Bank{
			"fjhxbc","cn","福建海峡银行","","","",
			[]int{621267,623063},
		},
	"dc", 18,
	}
	banks["cn_fjhxbc_pc_18"] = &BankPro{
		Bank{
			"fjhxbc","cn","福建海峡银行","","","",
			[]int{620043},
		},
	"pc", 18,
	}
	banks["cn_jlbank_dc_19"] = &BankPro{
		Bank{
			"jlbank","cn","吉林银行","","","",
			[]int{622865,623131},
		},
	"dc", 19,
	}
	banks["cn_jlbank_dc_16"] = &BankPro{
		Bank{
			"jlbank","cn","吉林银行","","","",
			[]int{940012},
		},
	"dc", 16,
	}
	banks["cn_jlbank_cc_16"] = &BankPro{
		Bank{
			"jlbank","cn","吉林银行","","","",
			[]int{622178,622179,628358},
		},
	"cc", 16,
	}
	banks["cn_hkb_dc_18"] = &BankPro{
		Bank{
			"hkb","cn","汉口银行","","","",
			[]int{990027},
		},
	"dc", 18,
	}
	banks["cn_hkb_dc_16"] = &BankPro{
		Bank{
			"hkb","cn","汉口银行","","","",
			[]int{622325,623105,623029},
		},
	"dc", 16,
	}
	banks["cn_sjbank_dc_18"] = &BankPro{
		Bank{
			"sjbank","cn","盛京银行","","","",
			[]int{566666},
		},
	"dc", 18,
	}
	banks["cn_sjbank_dc_19"] = &BankPro{
		Bank{
			"sjbank","cn","盛京银行","","","",
			[]int{622455,940039},
		},
	"dc", 19,
	}
	banks["cn_sjbank_dc_16"] = &BankPro{
		Bank{
			"sjbank","cn","盛京银行","","","",
			[]int{623108,623081},
		},
	"dc", 16,
	}
	banks["cn_sjbank_cc_16"] = &BankPro{
		Bank{
			"sjbank","cn","盛京银行","","","",
			[]int{622466,628285},
		},
	"cc", 16,
	}
	banks["cn_dlb_dc_17"] = &BankPro{
		Bank{
			"dlb","cn","大连银行","","","",
			[]int{603708},
		},
	"dc", 17,
	}
	banks["cn_dlb_dc_19"] = &BankPro{
		Bank{
			"dlb","cn","大连银行","","","",
			[]int{622993,623069,623070,623172,623173},
		},
	"dc", 19,
	}
	banks["cn_dlb_cc_16"] = &BankPro{
		Bank{
			"dlb","cn","大连银行","","","",
			[]int{622383,622385,628299},
		},
	"cc", 16,
	}
	banks["cn_bhb_dc_19"] = &BankPro{
		Bank{
			"bhb","cn","河北银行","","","",
			[]int{622498,622499,623000,940046},
		},
	"dc", 19,
	}
	banks["cn_bhb_cc_16"] = &BankPro{
		Bank{
			"bhb","cn","河北银行","","","",
			[]int{622921,628321},
		},
	"cc", 16,
	}
	banks["cn_urmqccb_dc_19"] = &BankPro{
		Bank{
			"urmqccb","cn","乌鲁木齐市商业银行","","","",
			[]int{621751,622143,940001,621754},
		},
	"dc", 19,
	}
	banks["cn_urmqccb_cc_16"] = &BankPro{
		Bank{
			"urmqccb","cn","乌鲁木齐市商业银行","","","",
			[]int{622476,628278},
		},
	"cc", 16,
	}
	banks["cn_sxcb_dc_16"] = &BankPro{
		Bank{
			"sxcb","cn","绍兴银行","","","",
			[]int{622486},
		},
	"dc", 16,
	}
	banks["cn_sxcb_dc_18"] = &BankPro{
		Bank{
			"sxcb","cn","绍兴银行","","","",
			[]int{603602,623026,623086},
		},
	"dc", 18,
	}
	banks["cn_sxcb_cc_16"] = &BankPro{
		Bank{
			"sxcb","cn","绍兴银行","","","",
			[]int{628291},
		},
	"cc", 16,
	}
	banks["cn_cdcb_dc_19"] = &BankPro{
		Bank{
			"cdcb","cn","成都商业银行","","","",
			[]int{622152,622154,622996,622997,940027,622153,622135,621482,621532},
		},
	"dc", 19,
	}
	banks["cn_fscb_dc_17"] = &BankPro{
		Bank{
			"fscb","cn","抚顺银行","","","",
			[]int{622442},
		},
	"dc", 17,
	}
	banks["cn_fscb_dc_18"] = &BankPro{
		Bank{
			"fscb","cn","抚顺银行","","","",
			[]int{940053},
		},
	"dc", 18,
	}
	banks["cn_fscb_dc_19"] = &BankPro{
		Bank{
			"fscb","cn","抚顺银行","","","",
			[]int{622442,623099},
		},
	"dc", 19,
	}
	banks["cn_zzbank_dc_19"] = &BankPro{
		Bank{
			"zzbank","cn","郑州银行","","","",
			[]int{622421},
		},
	"dc", 19,
	}
	banks["cn_zzbank_dc_17"] = &BankPro{
		Bank{
			"zzbank","cn","郑州银行","","","",
			[]int{940056},
		},
	"dc", 17,
	}
	banks["cn_zzbank_dc_16"] = &BankPro{
		Bank{
			"zzbank","cn","郑州银行","","","",
			[]int{96828},
		},
	"dc", 16,
	}
	banks["cn_nxbank_dc_19"] = &BankPro{
		Bank{
			"nxbank","cn","宁夏银行","","","",
			[]int{621529,622429,621417,623089,623200},
		},
	"dc", 19,
	}
	banks["cn_nxbank_cc_16"] = &BankPro{
		Bank{
			"nxbank","cn","宁夏银行","","","",
			[]int{628214,625529,622428},
		},
	"cc", 16,
	}
	banks["cn_cqbank_dc_16"] = &BankPro{
		Bank{
			"cqbank","cn","重庆银行","","","",
			[]int{9896},
		},
	"dc", 16,
	}
	banks["cn_cqbank_dc_16"] = &BankPro{
		Bank{
			"cqbank","cn","重庆银行","","","",
			[]int{622134,940018,623016},
		},
	"dc", 16,
	}
	banks["cn_hrbank_dc_19"] = &BankPro{
		Bank{
			"hrbank","cn","哈尔滨银行","","","",
			[]int{621577,622425},
		},
	"dc", 19,
	}
	banks["cn_hrbank_dc_18"] = &BankPro{
		Bank{
			"hrbank","cn","哈尔滨银行","","","",
			[]int{940049},
		},
	"dc", 18,
	}
	banks["cn_hrbank_dc_17"] = &BankPro{
		Bank{
			"hrbank","cn","哈尔滨银行","","","",
			[]int{622425},
		},
	"dc", 17,
	}
	banks["cn_lzyh_dc_16"] = &BankPro{
		Bank{
			"lzyh","cn","兰州银行","","","",
			[]int{622139,940040,628263},
		},
	"dc", 16,
	}
	banks["cn_lzyh_dc_19"] = &BankPro{
		Bank{
			"lzyh","cn","兰州银行","","","",
			[]int{621242,621538,621496},
		},
	"dc", 19,
	}
	banks["cn_qdccb_dc_16"] = &BankPro{
		Bank{
			"qdccb","cn","青岛银行","","","",
			[]int{621252,622146,940061,628239},
		},
	"dc", 16,
	}
	banks["cn_qdccb_dc_19"] = &BankPro{
		Bank{
			"qdccb","cn","青岛银行","","","",
			[]int{621419,623170},
		},
	"dc", 19,
	}
	banks["cn_qhdccb_dc_19"] = &BankPro{
		Bank{
			"qhdccb","cn","秦皇岛市商业银行","","","",
			[]int{62249802,94004602},
		},
	"dc", 19,
	}
	banks["cn_qhdccb_dc_19"] = &BankPro{
		Bank{
			"qhdccb","cn","秦皇岛市商业银行","","","",
			[]int{621237,623003},
		},
	"dc", 19,
	}
	banks["cn_boqh_dc_17"] = &BankPro{
		Bank{
			"boqh","cn","青海银行","","","",
			[]int{622310,940068},
		},
	"dc", 17,
	}
	banks["cn_boqh_cc_16"] = &BankPro{
		Bank{
			"boqh","cn","青海银行","","","",
			[]int{622817,628287,625959},
		},
	"cc", 16,
	}
	banks["cn_boqh_cc_16"] = &BankPro{
		Bank{
			"boqh","cn","青海银行","","","",
			[]int{62536601},
		},
	"cc", 16,
	}
	banks["cn_tzcb_dc_16"] = &BankPro{
		Bank{
			"tzcb","cn","台州银行","","","",
			[]int{622427},
		},
	"dc", 16,
	}
	banks["cn_tzcb_dc_17"] = &BankPro{
		Bank{
			"tzcb","cn","台州银行","","","",
			[]int{940069},
		},
	"dc", 17,
	}
	banks["cn_tzcb_dc_19"] = &BankPro{
		Bank{
			"tzcb","cn","台州银行","","","",
			[]int{623039},
		},
	"dc", 19,
	}
	banks["cn_tzcb_cc_16"] = &BankPro{
		Bank{
			"tzcb","cn","台州银行","","","",
			[]int{622321,628273},
		},
	"cc", 16,
	}
	banks["cn_tzcb_scc_16"] = &BankPro{
		Bank{
			"tzcb","cn","台州银行","","","",
			[]int{625001},
		},
	"scc", 16,
	}
	banks["cn_cscb_dc_18"] = &BankPro{
		Bank{
			"cscb","cn","长沙银行","","","",
			[]int{694301},
		},
	"dc", 18,
	}
	banks["cn_cscb_dc_19"] = &BankPro{
		Bank{
			"cscb","cn","长沙银行","","","",
			[]int{940071,622368,621446},
		},
	"dc", 19,
	}
	banks["cn_cscb_cc_16"] = &BankPro{
		Bank{
			"cscb","cn","长沙银行","","","",
			[]int{625901,622898,622900,628281,628282,622806,628283},
		},
	"cc", 16,
	}
	banks["cn_cscb_pc_19"] = &BankPro{
		Bank{
			"cscb","cn","长沙银行","","","",
			[]int{620519},
		},
	"pc", 19,
	}
	banks["cn_boqz_dc_18"] = &BankPro{
		Bank{
			"boqz","cn","泉州银行","","","",
			[]int{683970,940074},
		},
	"dc", 18,
	}
	banks["cn_boqz_dc_19"] = &BankPro{
		Bank{
			"boqz","cn","泉州银行","","","",
			[]int{622370},
		},
	"dc", 19,
	}
	banks["cn_boqz_dc_19"] = &BankPro{
		Bank{
			"boqz","cn","泉州银行","","","",
			[]int{621437},
		},
	"dc", 19,
	}
	banks["cn_boqz_cc_16"] = &BankPro{
		Bank{
			"boqz","cn","泉州银行","","","",
			[]int{628319},
		},
	"cc", 16,
	}
	banks["cn_bsb_dc_17"] = &BankPro{
		Bank{
			"bsb","cn","包商银行","","","",
			[]int{622336,621760},
		},
	"dc", 17,
	}
	banks["cn_bsb_dc_16"] = &BankPro{
		Bank{
			"bsb","cn","包商银行","","","",
			[]int{622165},
		},
	"dc", 16,
	}
	banks["cn_bsb_cc_16"] = &BankPro{
		Bank{
			"bsb","cn","包商银行","","","",
			[]int{622315,625950,628295},
		},
	"cc", 16,
	}
	banks["cn_daqingb_dc_19"] = &BankPro{
		Bank{
			"daqingb","cn","龙江银行","","","",
			[]int{621037,621097,621588,622977},
		},
	"dc", 19,
	}
	banks["cn_daqingb_dc_19"] = &BankPro{
		Bank{
			"daqingb","cn","龙江银行","","","",
			[]int{62321601},
		},
	"dc", 19,
	}
	banks["cn_daqingb_dc_16"] = &BankPro{
		Bank{
			"daqingb","cn","龙江银行","","","",
			[]int{622860},
		},
	"dc", 16,
	}
	banks["cn_daqingb_cc_16"] = &BankPro{
		Bank{
			"daqingb","cn","龙江银行","","","",
			[]int{622644,628333},
		},
	"cc", 16,
	}
	banks["cn_shrcb_dc_16"] = &BankPro{
		Bank{
			"shrcb","cn","上海农商银行","","","",
			[]int{622478,940013,621495},
		},
	"dc", 16,
	}
	banks["cn_shrcb_scc_16"] = &BankPro{
		Bank{
			"shrcb","cn","上海农商银行","","","",
			[]int{625500},
		},
	"scc", 16,
	}
	banks["cn_shrcb_cc_16"] = &BankPro{
		Bank{
			"shrcb","cn","上海农商银行","","","",
			[]int{622611,622722,628211,625989},
		},
	"cc", 16,
	}
	banks["cn_zjql_scc_16"] = &BankPro{
		Bank{
			"zjql","cn","浙江泰隆商业银行","","","",
			[]int{622717},
		},
	"scc", 16,
	}
	banks["cn_zjql_cc_16"] = &BankPro{
		Bank{
			"zjql","cn","浙江泰隆商业银行","","","",
			[]int{628275,622565,622287},
		},
	"cc", 16,
	}
	banks["cn_h3cb_dc_19"] = &BankPro{
		Bank{
			"h3cb","cn","内蒙古银行","","","",
			[]int{622147,621633},
		},
	"dc", 19,
	}
	banks["cn_h3cb_cc_16"] = &BankPro{
		Bank{
			"h3cb","cn","内蒙古银行","","","",
			[]int{628252},
		},
	"cc", 16,
	}
	banks["cn_bgb_dc_16"] = &BankPro{
		Bank{
			"bgb","cn","广西北部湾银行","","","",
			[]int{623001},
		},
	"dc", 16,
	}
	banks["cn_bgb_cc_16"] = &BankPro{
		Bank{
			"bgb","cn","广西北部湾银行","","","",
			[]int{628227},
		},
	"cc", 16,
	}
	banks["cn_glbank_dc_17"] = &BankPro{
		Bank{
			"glbank","cn","桂林银行","","","",
			[]int{621456},
		},
	"dc", 17,
	}
	banks["cn_glbank_dc_19"] = &BankPro{
		Bank{
			"glbank","cn","桂林银行","","","",
			[]int{621562},
		},
	"dc", 19,
	}
	banks["cn_glbank_cc_16"] = &BankPro{
		Bank{
			"glbank","cn","桂林银行","","","",
			[]int{628219},
		},
	"cc", 16,
	}
	banks["cn_daqingb_dc_19"] = &BankPro{
		Bank{
			"daqingb","cn","龙江银行","","","",
			[]int{621037,621097,621588,622977},
		},
	"dc", 19,
	}
	banks["cn_daqingb_dc_19"] = &BankPro{
		Bank{
			"daqingb","cn","龙江银行","","","",
			[]int{62321601},
		},
	"dc", 19,
	}
	banks["cn_daqingb_dc_16"] = &BankPro{
		Bank{
			"daqingb","cn","龙江银行","","","",
			[]int{622475,622860},
		},
	"dc", 16,
	}
	banks["cn_daqingb_scc_16"] = &BankPro{
		Bank{
			"daqingb","cn","龙江银行","","","",
			[]int{625588},
		},
	"scc", 16,
	}
	banks["cn_daqingb_cc_16"] = &BankPro{
		Bank{
			"daqingb","cn","龙江银行","","","",
			[]int{622270,628368,625090,622644,628333},
		},
	"cc", 16,
	}
	banks["cn_cdrcb_dc_19"] = &BankPro{
		Bank{
			"cdrcb","cn","成都农村商业银行","","","",
			[]int{623088},
		},
	"dc", 19,
	}
	banks["cn_cdrcb_cc_16"] = &BankPro{
		Bank{
			"cdrcb","cn","成都农村商业银行","","","",
			[]int{622829,628301,622808,628308},
		},
	"cc", 16,
	}
	banks["cn_fjnx_dc_19"] = &BankPro{
		Bank{
			"fjnx","cn","福建省农村信用社联合社","","","",
			[]int{622127,622184,621701,621251,621589,623036},
		},
	"dc", 19,
	}
	banks["cn_fjnx_cc_16"] = &BankPro{
		Bank{
			"fjnx","cn","福建省农村信用社联合社","","","",
			[]int{628232,622802,622290},
		},
	"cc", 16,
	}
	banks["cn_trcb_dc_19"] = &BankPro{
		Bank{
			"trcb","cn","天津农村商业银行","","","",
			[]int{622531,622329},
		},
	"dc", 19,
	}
	banks["cn_trcb_cc_16"] = &BankPro{
		Bank{
			"trcb","cn","天津农村商业银行","","","",
			[]int{622829,628301},
		},
	"cc", 16,
	}
	banks["cn_jsrcu_dc_19"] = &BankPro{
		Bank{
			"jsrcu","cn","江苏省农村信用社联合社","","","",
			[]int{621578,623066,622452,622324},
		},
	"dc", 19,
	}
	banks["cn_jsrcu_cc_16"] = &BankPro{
		Bank{
			"jsrcu","cn","江苏省农村信用社联合社","","","",
			[]int{622815,622816,628226},
		},
	"cc", 16,
	}
	banks["cn_slh_cc_16"] = &BankPro{
		Bank{
			"slh","cn","湖南农村信用社联合社","","","",
			[]int{622906,628386,625519,625506},
		},
	"cc", 16,
	}
	banks["cn_jxncx_dc_16"] = &BankPro{
		Bank{
			"jxncx","cn","江西省农村信用社联合社","","","",
			[]int{621592},
		},
	"dc", 16,
	}
	banks["cn_jxncx_cc_16"] = &BankPro{
		Bank{
			"jxncx","cn","江西省农村信用社联合社","","","",
			[]int{628392},
		},
	"cc", 16,
	}
	banks["cn_scbbank_dc_19"] = &BankPro{
		Bank{
			"scbbank","cn","商丘市商业银行","","","",
			[]int{621748},
		},
	"dc", 19,
	}
	banks["cn_scbbank_cc_16"] = &BankPro{
		Bank{
			"scbbank","cn","商丘市商业银行","","","",
			[]int{628271},
		},
	"cc", 16,
	}
	banks["cn_hrxjb_dc_19"] = &BankPro{
		Bank{
			"hrxjb","cn","华融湘江银行","","","",
			[]int{621366,621388},
		},
	"dc", 19,
	}
	banks["cn_hrxjb_cc_16"] = &BankPro{
		Bank{
			"hrxjb","cn","华融湘江银行","","","",
			[]int{628328},
		},
	"cc", 16,
	}
	banks["cn_hsbk_dc_19"] = &BankPro{
		Bank{
			"hsbk","cn","衡水市商业银行","","","",
			[]int{621239,623068},
		},
	"dc", 19,
	}
	banks["cn_cqncsycz_dc_19"] = &BankPro{
		Bank{
			"cqncsycz","cn","重庆南川石银村镇银行","","","",
			[]int{621653004},
		},
	"dc", 19,
	}
	banks["cn_hnrcc_dc_19"] = &BankPro{
		Bank{
			"hnrcc","cn","湖南省农村信用社联合社","","","",
			[]int{622169,621519,621539,623090},
		},
	"dc", 19,
	}
	banks["cn_xtb_dc_19"] = &BankPro{
		Bank{
			"xtb","cn","邢台银行","","","",
			[]int{621238,620528},
		},
	"dc", 19,
	}
	banks["cn_lprdncxys_cc_16"] = &BankPro{
		Bank{
			"lprdncxys","cn","临汾市尧都区农村信用合作联社","","","",
			[]int{628382,625158},
		},
	"cc", 16,
	}
	banks["cn_dyccb_dc_18"] = &BankPro{
		Bank{
			"dyccb","cn","东营银行","","","",
			[]int{621004},
		},
	"dc", 18,
	}
	banks["cn_dyccb_cc_16"] = &BankPro{
		Bank{
			"dyccb","cn","东营银行","","","",
			[]int{628217},
		},
	"cc", 16,
	}
	banks["cn_srbank_dc_16"] = &BankPro{
		Bank{
			"srbank","cn","上饶银行","","","",
			[]int{621416},
		},
	"dc", 16,
	}
	banks["cn_srbank_cc_16"] = &BankPro{
		Bank{
			"srbank","cn","上饶银行","","","",
			[]int{628217},
		},
	"cc", 16,
	}
	banks["cn_dzbank_dc_19"] = &BankPro{
		Bank{
			"dzbank","cn","德州银行","","","",
			[]int{622937},
		},
	"dc", 19,
	}
	banks["cn_dzbank_cc_16"] = &BankPro{
		Bank{
			"dzbank","cn","德州银行","","","",
			[]int{628397},
		},
	"cc", 16,
	}
	banks["cn_cdb_cc_16"] = &BankPro{
		Bank{
			"cdb","cn","承德银行","","","",
			[]int{628229},
		},
	"cc", 16,
	}
	banks["cn_ynrcc_cc_16"] = &BankPro{
		Bank{
			"ynrcc","cn","云南省农村信用社","","","",
			[]int{622469,628307},
		},
	"cc", 16,
	}
	banks["cn_lzccb_dc_18"] = &BankPro{
		Bank{
			"lzccb","cn","柳州银行","","","",
			[]int{622292,622291,621412},
		},
	"dc", 18,
	}
	banks["cn_lzccb_dc_16"] = &BankPro{
		Bank{
			"lzccb","cn","柳州银行","","","",
			[]int{622880,622881},
		},
	"dc", 16,
	}
	banks["cn_lzccb_cc_15"] = &BankPro{
		Bank{
			"lzccb","cn","柳州银行","","","",
			[]int{62829},
		},
	"cc", 15,
	}
	banks["cn_whsybank_dc_16"] = &BankPro{
		Bank{
			"whsybank","cn","威海市商业银行","","","",
			[]int{623102},
		},
	"dc", 16,
	}
	banks["cn_whsybank_cc_16"] = &BankPro{
		Bank{
			"whsybank","cn","威海市商业银行","","","",
			[]int{628234},
		},
	"cc", 16,
	}
	banks["cn_hzbank_cc_16"] = &BankPro{
		Bank{
			"hzbank","cn","湖州银行","","","",
			[]int{628306},
		},
	"cc", 16,
	}
	banks["cn_bankwf_dc_16"] = &BankPro{
		Bank{
			"bankwf","cn","潍坊银行","","","",
			[]int{622391,940072},
		},
	"dc", 16,
	}
	banks["cn_bankwf_cc_16"] = &BankPro{
		Bank{
			"bankwf","cn","潍坊银行","","","",
			[]int{628391},
		},
	"cc", 16,
	}
	banks["cn_gzb_dc_19"] = &BankPro{
		Bank{
			"gzb","cn","赣州银行","","","",
			[]int{622967,940073},
		},
	"dc", 19,
	}
	banks["cn_gzb_cc_16"] = &BankPro{
		Bank{
			"gzb","cn","赣州银行","","","",
			[]int{628233},
		},
	"cc", 16,
	}
	banks["cn_rzgwybank_cc_16"] = &BankPro{
		Bank{
			"rzgwybank","cn","日照银行","","","",
			[]int{628257},
		},
	"cc", 16,
	}
	banks["cn_ncb_dc_16"] = &BankPro{
		Bank{
			"ncb","cn","南昌银行","","","",
			[]int{621269,622275},
		},
	"dc", 16,
	}
	banks["cn_ncb_dc_17"] = &BankPro{
		Bank{
			"ncb","cn","南昌银行","","","",
			[]int{940006},
		},
	"dc", 17,
	}
	banks["cn_ncb_cc_17"] = &BankPro{
		Bank{
			"ncb","cn","南昌银行","","","",
			[]int{628305},
		},
	"cc", 17,
	}
	banks["cn_gycb_dc_19"] = &BankPro{
		Bank{
			"gycb","cn","贵阳银行","","","",
			[]int{622133,621735},
		},
	"dc", 19,
	}
	banks["cn_gycb_dc_16"] = &BankPro{
		Bank{
			"gycb","cn","贵阳银行","","","",
			[]int{888},
		},
	"dc", 16,
	}
	banks["cn_gycb_cc_16"] = &BankPro{
		Bank{
			"gycb","cn","贵阳银行","","","",
			[]int{628213},
		},
	"cc", 16,
	}
	banks["cn_bojz_dc_17"] = &BankPro{
		Bank{
			"bojz","cn","锦州银行","","","",
			[]int{622990,940003},
		},
	"dc", 17,
	}
	banks["cn_bojz_cc_16"] = &BankPro{
		Bank{
			"bojz","cn","锦州银行","","","",
			[]int{628261},
		},
	"cc", 16,
	}
	banks["cn_qsbank_dc_17"] = &BankPro{
		Bank{
			"qsbank","cn","齐商银行","","","",
			[]int{622311,940057},
		},
	"dc", 17,
	}
	banks["cn_qsbank_cc_16"] = &BankPro{
		Bank{
			"qsbank","cn","齐商银行","","","",
			[]int{628311},
		},
	"cc", 16,
	}
	banks["cn_rboz_dc_19"] = &BankPro{
		Bank{
			"rboz","cn","珠海华润银行","","","",
			[]int{622363,940048},
		},
	"dc", 19,
	}
	banks["cn_rboz_cc_16"] = &BankPro{
		Bank{
			"rboz","cn","珠海华润银行","","","",
			[]int{628270},
		},
	"cc", 16,
	}
	banks["cn_hldccb_dc_16"] = &BankPro{
		Bank{
			"hldccb","cn","葫芦岛市商业银行","","","",
			[]int{622398,940054},
		},
	"dc", 16,
	}
	banks["cn_hbc_dc_17"] = &BankPro{
		Bank{
			"hbc","cn","宜昌市商业银行","","","",
			[]int{940055},
		},
	"dc", 17,
	}
	banks["cn_hbc_cc_17"] = &BankPro{
		Bank{
			"hbc","cn","宜昌市商业银行","","","",
			[]int{622397},
		},
	"cc", 17,
	}
	banks["cn_hzcb_dc_18"] = &BankPro{
		Bank{
			"hzcb","cn","杭州商业银行","","","",
			[]int{603367,622878},
		},
	"dc", 18,
	}
	banks["cn_hzcb_cc_17"] = &BankPro{
		Bank{
			"hzcb","cn","杭州商业银行","","","",
			[]int{622397},
		},
	"cc", 17,
	}
	banks["cn_jsbank_dc_19"] = &BankPro{
		Bank{
			"jsbank","cn","苏州市商业银行","","","",
			[]int{603506},
		},
	"dc", 19,
	}
	banks["cn_lycb_dc_17"] = &BankPro{
		Bank{
			"lycb","cn","辽阳银行","","","",
			[]int{622399,940043},
		},
	"dc", 17,
	}
	banks["cn_lyb_dc_17"] = &BankPro{
		Bank{
			"lyb","cn","洛阳银行","","","",
			[]int{622420,940041},
		},
	"dc", 17,
	}
	banks["cn_jzcbank_dc_19"] = &BankPro{
		Bank{
			"jzcbank","cn","焦作市商业银行","","","",
			[]int{622338},
		},
	"dc", 19,
	}
	banks["cn_jzcbank_dc_16"] = &BankPro{
		Bank{
			"jzcbank","cn","焦作市商业银行","","","",
			[]int{940032},
		},
	"dc", 16,
	}
	banks["cn_zjccb_dc_16"] = &BankPro{
		Bank{
			"zjccb","cn","镇江市商业银行","","","",
			[]int{622394,940025},
		},
	"dc", 16,
	}
	banks["cn_fgxybank_dc_16"] = &BankPro{
		Bank{
			"fgxybank","cn","法国兴业银行","","","",
			[]int{621245},
		},
	"dc", 16,
	}
	banks["cn_dybank_dc_19"] = &BankPro{
		Bank{
			"dybank","cn","大华银行","","","",
			[]int{621328},
		},
	"dc", 19,
	}
	banks["cn_diyebank_dc_19"] = &BankPro{
		Bank{
			"diyebank","cn","企业银行","","","",
			[]int{621651},
		},
	"dc", 19,
	}
	banks["cn_hqbank_dc_16"] = &BankPro{
		Bank{
			"hqbank","cn","华侨银行","","","",
			[]int{621077},
		},
	"dc", 16,
	}
	banks["cn_hsb_dc_19"] = &BankPro{
		Bank{
			"hsb","cn","恒生银行","","","",
			[]int{622409,621441},
		},
	"dc", 19,
	}
	banks["cn_hsb_dc_17"] = &BankPro{
		Bank{
			"hsb","cn","恒生银行","","","",
			[]int{622410,621440},
		},
	"dc", 17,
	}
	banks["cn_hsb_dc_16"] = &BankPro{
		Bank{
			"hsb","cn","恒生银行","","","",
			[]int{622950,622951},
		},
	"dc", 16,
	}
	banks["cn_hsb_cc_16"] = &BankPro{
		Bank{
			"hsb","cn","恒生银行","","","",
			[]int{625026,625024,622376,622378,622377,625092},
		},
	"cc", 16,
	}
	banks["cn_lsb_dc_19"] = &BankPro{
		Bank{
			"lsb","cn","临沂商业银行","","","",
			[]int{622359,940066},
		},
	"dc", 19,
	}
	banks["cn_ytcb_dc_16"] = &BankPro{
		Bank{
			"ytcb","cn","烟台商业银行","","","",
			[]int{622886},
		},
	"dc", 16,
	}
	banks["cn_qlb_dc_19"] = &BankPro{
		Bank{
			"qlb","cn","齐鲁银行","","","",
			[]int{940008,622379},
		},
	"dc", 19,
	}
	banks["cn_qlb_cc_16"] = &BankPro{
		Bank{
			"qlb","cn","齐鲁银行","","","",
			[]int{628379},
		},
	"cc", 16,
	}
	banks["cn_bccc_dc_16"] = &BankPro{
		Bank{
			"bccc","cn","BC卡公司","","","",
			[]int{620011,620027,620031,620039,620103,620106,620120,620123,620125,620220,620278,620812,621006,621011,621012,621020,621023,621025,621027,621031,620132,621039,621078,621220,621003},
		},
	"dc", 16,
	}
	banks["cn_bccc_cc_16"] = &BankPro{
		Bank{
			"bccc","cn","BC卡公司","","","",
			[]int{625003,625011,625012,625020,625023,625025,625027,625031,621032,625039,625078,625079,625103,625106,625006,625112,625120,625123,625125,625127,625131,625032,625139,625178,625179,625220,625320,625111,625132,625244},
		},
	"cc", 16,
	}
	banks["cn_cyb_dc_16"] = &BankPro{
		Bank{
			"cyb","cn","集友银行","","","",
			[]int{622355,623042},
		},
	"dc", 16,
	}
	banks["cn_cyb_dc_19"] = &BankPro{
		Bank{
			"cyb","cn","集友银行","","","",
			[]int{621043,621742},
		},
	"dc", 19,
	}
	banks["cn_cyb_cc_16"] = &BankPro{
		Bank{
			"cyb","cn","集友银行","","","",
			[]int{622352,622353,625048,625053,625060},
		},
	"cc", 16,
	}
	banks["cn_cyb_pc_16"] = &BankPro{
		Bank{
			"cyb","cn","集友银行","","","",
			[]int{620206,620207},
		},
	"pc", 16,
	}
	banks["cn_tfb_dc_19"] = &BankPro{
		Bank{
			"tfb","cn","大丰银行","","","",
			[]int{622547,622548,622546},
		},
	"dc", 19,
	}
	banks["cn_tfb_cc_16"] = &BankPro{
		Bank{
			"tfb","cn","大丰银行","","","",
			[]int{625198,625196,625147},
		},
	"cc", 16,
	}
	banks["cn_tfb_pc_19"] = &BankPro{
		Bank{
			"tfb","cn","大丰银行","","","",
			[]int{620072},
		},
	"pc", 19,
	}
	banks["cn_tfb_pc_16"] = &BankPro{
		Bank{
			"tfb","cn","大丰银行","","","",
			[]int{620204,620205},
		},
	"pc", 16,
	}
	banks["cn_aeon_dc_16"] = &BankPro{
		Bank{
			"aeon","cn","AEON信贷财务亚洲有限公司","","","",
			[]int{621064,622941,622974},
		},
	"dc", 16,
	}
	banks["cn_aeon_cc_16"] = &BankPro{
		Bank{
			"aeon","cn","AEON信贷财务亚洲有限公司","","","",
			[]int{622493},
		},
	"cc", 16,
	}
	banks["cn_mabda_dc_19"] = &BankPro{
		Bank{
			"mabda","cn","澳门BDA","","","",
			[]int{621274,621324},
		},
	"dc", 19,
	}
	BanksProMap["cn"] = banks
	MaxLenProMap["cn"] = 10
}