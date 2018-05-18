package bankcard

func init() {
	Banks["ar_galicia"] = &Bank{
		"galicia","ar","Banco Galicia/Naranja","Galicia Bank / Naranja","http://www.bancogalicia.com/","#ff6319",
		[]int{402918},
	}
	Banks["ar_hsbc"] = &Bank{
		"hsbc","ar","HSBC","HSBC","https://www.hsbc.com.ar/es/","#dd1f26",
		[]int{450910,525489},
	}
	Banks["ar_macro"] = &Bank{
		"macro","ar","Banco Macro","Macro Bank","https://www.macro.com.ar","#0c75a5",
		[]int{451764},
	}
	Banks["ar_santanderrio"] = &Bank{
		"santanderrio","ar","Santander Rio","Santander Rio","http://santanderrio.com.ar/","#ef3340",
		[]int{405071},
	}
	Banks["au_anz"] = &Bank{
		"anz","au","ANZ","ANZ","https://www.anz.com/","#007dba",
		[]int{407220,450949,456462,456468,456469,546827},
	}
	Banks["au_bwe"] = &Bank{
		"bwe","au","Bankwest","Bankwest","http://www.bankwest.com.au/","#f79534",
		[]int{456444,522980,528013,532737,540482,543568},
	}
	Banks["au_cba"] = &Bank{
		"cba","au","Commonwealth Bank of Australia","Commonwealth Bank of Australia","https://www.commbank.com.au/","#f3c40e",
		[]int{405221,494052,494053,521729,535316,535318,535319,552033,555005},
	}
	Banks["au_go"] = &Bank{
		"go","au","GO","GO","https://www.gomastercard.com.au/","#486d11",
		[]int{521893},
	}
	Banks["au_ing"] = &Bank{
		"ing","au","ING Direct","ING Direct","https://www.ingdirect.com.au/","#ff6200",
		[]int{462263},
	}
	Banks["au_lom"] = &Bank{
		"lom","au","Lombard","Lombard","https://www.lombardfinance.com.au/","#002d55",
		[]int{445013},
	}
	Banks["au_suncorp"] = &Bank{
		"suncorp","au","Suncorp","Suncorp","http://www.suncorp.com.au","#006f66",
		[]int{439239,456432},
	}
	Banks["be_amexbelgiumcorporate"] = &Bank{
		"amexbelgiumcorporate","be","American Express","American Express","https://network.americanexpress.com/be/fr/homepage/landing.aspx","#67b198",
		[]int{374190},
	}
	Banks["be_belfius"] = &Bank{
		"belfius","be","Belfius Bank","Belfius Bank","https://www.belfius.be","#325c6c",
		[]int{670305},
	}
	Banks["be_ingbelgium"] = &Bank{
		"ingbelgium","be","ING Belgium","ING Belgium","https://www.ing.be/","#ff6600",
		[]int{42360},
	}
	Banks["by_absolutbank"] = &Bank{
		"absolutbank","by","Абсолютбанк","Absolutbank","http://www.absolutbank.by","#b72930",
		[]int{458563},
	}
	Banks["by_alfabank"] = &Bank{
		"alfabank","by","Альфабанк","Alfabank","https://www.alfabank.by/","#ee3124",
		[]int{437350,458520,458521,458522,458523},
	}
	Banks["by_belapb"] = &Bank{
		"belapb","by","Белагропромбанк","Belagroprombank","http://www.belapb.by","#fab606",
		[]int{432036,432037,432038,432039,432040,432041,520195,522208,532308,547232},
	}
	Banks["by_belarusbank"] = &Bank{
		"belarusbank","by","Беларусбанк","Belarusbank","http://www.belarusbank.by","#127d49",
		[]int{401803,409087,424641,425519,425520,425521,510035,543044,543634,544598,544650,547087,547091,547232,547504,557075,557076},
	}
	Banks["by_belaruskynarodny"] = &Bank{
		"belaruskynarodny","by","Белорусский народный банк","Beloruskii Narodny Bank","http://www.bnb.by","#f89828",
		[]int{450040,450041,450042,489062},
	}
	Banks["by_belgazprom"] = &Bank{
		"belgazprom","by","Белгазпромбанк","Belgazprombank","http://www.belgazprombank.by","#0079c2",
		[]int{428620,428621,428622,428623,428624,518597,518707,529922},
	}
	Banks["by_belinvestbank"] = &Bank{
		"belinvestbank","by","Белинвестбанк","Belinvestbank","http://www.belinvestbank.by/","#009881",
		[]int{401346,434484,434485,434486,434487,458291,547691,557749,557884,557885},
	}
	Banks["by_belswissbank"] = &Bank{
		"belswissbank","by","БелСвиссБанк","BelSwiss","https://www.bsb.by/","#ee1d23",
		[]int{410745,410746},
	}
	Banks["by_belveb"] = &Bank{
		"belveb","by","Банк БелВЭБ","BelVEB Bank","http://www.bveb.by/","#f37637",
		[]int{414088,414297,414298,422887,447514,447515,470469,499835,520658,537000,544206,544578,544650,547747,548475},
	}
	Banks["by_bpssberbank"] = &Bank{
		"bpssberbank","by","БПС-Сбербанк","BPS-Sberbank","http://www.bps-sberbank.by/","#1b991b",
		[]int{413856,413857,431836,431837,431838,437449,540234,544454,558682},
	}
	Banks["by_deltabank"] = &Bank{
		"deltabank","by","Дельтабанк","Delta Bank","http://deltabank.by/","#d81e05",
		[]int{467939,467940,467941},
	}
	Banks["by_homecredit"] = &Bank{
		"homecredit","by","Хоум Кредит Банк","Home Credit Bank","http://homecredit.by/","#e03a3e",
		[]int{510415,510428,528065},
	}
	Banks["by_moskvaminsk"] = &Bank{
		"moskvaminsk","by","Банк Москва-Минск","Bank Moskva-Minsk","http://www.mmbank.by/","#f2000d",
		[]int{446505},
	}
	Banks["by_mtb"] = &Bank{
		"mtb","by","МТБанк","MTBank","http://www.mtbank.by","#005aab",
		[]int{417753,417754,417755,464990},
	}
	Banks["by_paritetbank"] = &Bank{
		"paritetbank","by","Паритетбанк","Paritetbank","http://www.paritetbank.by/","#0360a7",
		[]int{416676,416677,427870,429737,437391,459056},
	}
	Banks["by_priorbank"] = &Bank{
		"priorbank","by","Приорбанк","Priorbank","http://www.priorbank.by","#ffff00",
		[]int{413059,416064,417770,437458,445987,458285,471326,472950,480949,491698,542645,548952,558349},
	}
	Banks["by_trustbank"] = &Bank{
		"trustbank","by","Трастбанк","Trustbank","http://www.trustbank.by/","#a81429",
		[]int{411856,432231,432232},
	}
	Banks["by_vtb"] = &Bank{
		"vtb","by","ВТБ (Беларусь)","VTB (Belarus)","http://www.vtb-bank.by","#002e77",
		[]int{445451,445452,445453,445454,471346},
	}
	Banks["ca_scotiabank"] = &Bank{
		"scotiabank","ca","Scotiabank","Scotiabank","http://www.scotiabank.com/","#d81e05",
		[]int{453600},
	}
	Banks["ca_tangerine"] = &Bank{
		"tangerine","ca","Tangerine","Tangerine","https://www.tangerine.ca/","#fa823c",
		[]int{601725},
	}
	Banks["ca_tdcanadatrust"] = &Bank{
		"tdcanadatrust","ca","TD Canada Trust","TD Canada Trust","https://www.tdcanadatrust.com","#20b14a",
		[]int{472409},
	}
	Banks["cn_abc"] = &Bank{
		"abc","cn","中国农业银行","Agricultural Bank of China","http://www.abchina.com/","#319c8b",
		[]int{403361,404117,404118,404119,404120,404121,49102,519412,519413,520082,520083,53591,552599,558730,622836,622837,622848},
	}
	Banks["cn_boc"] = &Bank{
		"boc","cn","中国银行","Bank of China","http://www.boc.cn/","#a71e32",
		[]int{409666,438088,622752,622760},
	}
	Banks["cn_boco"] = &Bank{
		"boco","cn","交通银行","Bank of Communications","http://www.bankcomm.com","#023978",
		[]int{458123,521899,622260},
	}
	Banks["cn_ccb"] = &Bank{
		"ccb","cn","中国建设银行","China Construction Bank","http://www.ccb.com/","#0c69b0",
		[]int{436742,436745,621700,622280},
	}
	Banks["cn_ceb"] = &Bank{
		"ceb","cn","中国光大银行","China Everbright Bank","http://www.cebbank.cn/","#691585",
		[]int{356837,356838,356839,356840,406252,406254,425862,481699,486497,543159,622161,622650,622655,622658,628201,628202},
	}
	Banks["cn_cgb"] = &Bank{
		"cgb","cn","中国银行","China Guangfa Bank","http://www.cgbchina.com.cn/","#e60021",
		[]int{520152,520382,548844,622568,911121},
	}
	Banks["cn_cib"] = &Bank{
		"cib","cn","兴业银行","Industrial Bank Co. Ltd.","http://www.cib.com.cn/","#004097",
		[]int{451289,451290,486493,486494,486861,523036,524070,527414,622901,622902,622922},
	}
	Banks["cn_citic"] = &Bank{
		"citic","cn","中信银行","China Citic Bank","http://bank.ecitic.com/","#c8151c",
		[]int{376966,376968,376969,400360,403391,403392,403393,404157,404158,404159,404171,404172,404173,404174,433666,433667,433668,433669,514906,518212,520108,556617,558916,622680,622688,622689,622690,622916,622918,622919,628206,628208,628209},
	}
	Banks["cn_cmb"] = &Bank{
		"cmb","cn","招商银行","China Merchants Bank","http://www.cmbchina.com","#b0110c",
		[]int{356885,356886,356889,356890,439188,439225,439226,439227,439228,439229,479229,518710,518718,521302,545619,545623,545947,545948,552534,552580,552581,552582,552583,552584,552585,552586,552588,552589,622575,622576,622577,622578,622579,622581,622588,628262,628362,645621},
	}
	Banks["cn_cmib"] = &Bank{
		"cmib","cn","中国民生银行","China Minsheng Bank","http://www.cmbc.com.cn/","#0053b6",
		[]int{407405,415599,421869,421870,421871,512466,517636,528948,552288,556610,622600,622601,622602,622603,622622,628258},
	}
	Banks["cn_hxb"] = &Bank{
		"hxb","cn","华夏银行","Huaxia Bank","http://www.hxb.com.cn/","#d80c18",
		[]int{523959,528708,528709,539867,539868,622636,622637},
	}
	Banks["cn_icbc"] = &Bank{
		"icbc","cn","中国工商银行","Industrial and Commercial Bank of China","http://www.icbc.com.cn","#c90000",
		[]int{427020,427030,530990,622200,622210,622215,622230,622235,955880},
	}
	Banks["cn_pab"] = &Bank{
		"pab","cn","平安银行","Ping An Bank","http://bank.pingan.com/","#e95303",
		[]int{435744,435745,526855,528020,622155,622156,622525,622526,622538,998801,998802},
	}
	Banks["cn_spdb"] = &Bank{
		"spdb","cn","浦东发展银行","Shanghai Pudong Development Bank","http://www.spdb.com.cn","#072154",
		[]int{622516,622517,622518,622519,622521,622522,622523},
	}
	Banks["cy_hellenicbank"] = &Bank{
		"hellenicbank","cy","Ελληνική Τράπεζα","Hellenic Bank","https://www.hellenicbank.com/","#003f7e",
		[]int{456747},
	}
	Banks["cz_airbank"] = &Bank{
		"airbank","cz","Airbank","Airbank","https://airbank.cz","#99cc33",
		[]int{516844},
	}
	Banks["cz_csob"] = &Bank{
		"csob","cz","ČSOB","ČSOB","https://www.csob.cz/","#0099cc",
		[]int{440753,543463},
	}
	Banks["cz_fio-banka"] = &Bank{
		"fio-banka","cz","Fio banka","Fio banka","https://www.fio.cz/","#00458a",
		[]int{510658},
	}
	Banks["de_n26"] = &Bank{
		"n26","de","N26","N26","https://n26.com/","#383838",
		[]int{535584,535585},
	}
	Banks["de_postbank"] = &Bank{
		"postbank","de","Postbank","Postbank","https://www.postbank.de/","#ffcc00",
		[]int{494116},
	}
	Banks["dk_jyskebank"] = &Bank{
		"jyskebank","dk","Jyske Bank","Jyske Bank","https://www.jyskebank.dk/","#005231",
		[]int{457172},
	}
	Banks["dk_nykredit"] = &Bank{
		"nykredit","dk","Nykredit","Nykredit","http://nykredit.dk/","#10137c",
		[]int{457181},
	}
	Banks["fi_nordeafinland"] = &Bank{
		"nordeafinland","fi","Nordea Pankki Suomi Oyj","Nordea","http://www.nordea.fi/","#055b90",
		[]int{492010},
	}
	Banks["fr_boursorama"] = &Bank{
		"boursorama","fr","Boursorama","Boursorama","http://www.boursorama.com/","#d20073",
		[]int{497958},
	}
	Banks["fr_caisseepargne"] = &Bank{
		"caisseepargne","fr","Caisse d'Épargne","Caisse d'Épargne","https://www.caisse-epargne.fr/","#e2001a",
		[]int{497874},
	}
	Banks["fr_creditmutuel"] = &Bank{
		"creditmutuel","fr","Credit Mutuel","Credit Mutuel","https://www.creditmutuel.com/","#11519b",
		[]int{513616,513742},
	}
	Banks["fr_ingdirect"] = &Bank{
		"ingdirect","fr","ING Direct","ING Direct","http://www.ingdirect.fr/","#ff6600",
		[]int{513625},
	}
	Banks["fr_lcl"] = &Bank{
		"lcl","fr","LCL","LCL","https://particuliers.secure.lcl.fr/","#263f92",
		[]int{497203},
	}
	Banks["fr_societe-generale"] = &Bank{
		"societe-generale","fr","Société Générale","Societe Generale","http://societegenerale.fr/","#e5022f",
		[]int{49735,497355},
	}
	Banks["gb_barclaycard"] = &Bank{
		"barclaycard","gb","Barclaycard","Barclaycard","http://www.barclaycard.co.uk/","#018fd0",
		[]int{492910},
	}
	Banks["gb_barclays"] = &Bank{
		"barclays","gb","Barclays Bank plc","Barclays Bank plc","https://www.barclays.co.uk/","#00adef",
		[]int{400115,408367,409400,409401,409402,430532,453978,453979,456725,465858,465859,465860,465861,465862,465863,465864,465865,465866,465867,465901,465902,465911,465921,465922,465923,484419,484420,491748,491749,491750,492902,492912},
	}
	Banks["gb_curve"] = &Bank{
		"curve","gb","Curve","Curve","https://www.imaginecurve.com","#0f2d4f",
		[]int{516273},
	}
	Banks["gb_halifax"] = &Bank{
		"halifax","gb","Halifax","Halifax","http://www.halifax.co.uk/","#0040bb",
		[]int{446238,446291,528683,543429,552073},
	}
	Banks["gb_hsbc"] = &Bank{
		"hsbc","gb","HSBC Bank plc","HSBC Bank plc","https://www.hsbc.co.uk/","#ee3226",
		[]int{465942,465943,465944,486483,543460},
	}
	Banks["gb_lloydstsb"] = &Bank{
		"lloydstsb","gb","Lloyds Bank plc","Lloyds Bank plc","https://www.lloydsbank.com/","#006a4d",
		[]int{476363,491735},
	}
	Banks["gb_metrobank"] = &Bank{
		"metrobank","gb","Metro Bank","Metro Bank","https://www.metrobankonline.co.uk","#013760",
		[]int{557361},
	}
	Banks["gb_monzo"] = &Bank{
		"monzo","gb","Monzo","Monzo","https://monzo.com","#f93845",
		[]int{535420},
	}
	Banks["gb_nationwide"] = &Bank{
		"nationwide","gb","Nationwide","Nationwide","http://www.nationwide.co.uk","#002878",
		[]int{454313,465935,475141,489396},
	}
	Banks["gb_natwest"] = &Bank{
		"natwest","gb","NatWest","NatWest","http://personal.natwest.com/personal.html","#42145f",
		[]int{426326,475117,475128,475129,475131,490743,490946,493600,493601,493602,493603,493604,493605,493606,493607,493608,493609,493610,493611,493612,493613,493614,493615,493616,493617,493618,493619,493620,493621,493622,493623,493624,493625,493626,493627,493628,493629,493630,493631,493632,493633,493634,493635,493636,493637,493638,493639,493640,493641,493642,493643,493644,493645,493646,493647,493648,493649,493650,493651,493652,493653,493655,493657,493658,493659,493661,493662,493663,493664,493665,493666,493667,493668,493669,493670,493671,493672,493673,493674,493675,493676,493677,493678,493679,493680,493681,493682,493683,493684,493685,493686,493687,493689,493690,493691,493692,493693,493694,493695,493696,493697,543166,545107,545108,545110,545112,545113,545121,545130,545135,545145,545149,545152,545154,545175,545176,545187,545188,545191,545461,545475,545486,545501,545513,545514,545515,545516,545517,545518,545520,545521,545522,545523,545524,552211,552217,552222,552230,552242,552249,552270,552272,552276,552282,552294,552295,552296,552298,675958,675961,675962,675977,675978,675998,676215,676701,676709,676718,676750,676751,676752,676753,676754,676755,676756,676757,676758,676759,676760,676761,676762,676798,676819,677121},
	}
	Banks["gb_revolut"] = &Bank{
		"revolut","gb","Revolut","Revolut","https://revolut.com/","#1574e2",
		[]int{539123},
	}
	Banks["gb_santander"] = &Bank{
		"santander","gb","Santander UK plc","Santander UK plc","http://www.santander.co.uk/","#ec0000",
		[]int{454742,475714,528689},
	}
	Banks["gb_soldo"] = &Bank{
		"soldo","gb","Soldo","Soldo","https://www.soldo.com","#fd8a43",
		[]int{557429},
	}
	Banks["gb_tsb"] = &Bank{
		"tsb","gb","TSB","TSB","http://tsb.co.uk/","#0098cf",
		[]int{476367,476368},
	}
	Banks["ie_aib"] = &Bank{
		"aib","ie","AIB Bank","AIB Bank","https://aib.ie","#7f2b7b",
		[]int{402344,402345,402346,402347,402348,402349,403023,406410,406411,406412,408531,408569,408570,409810,414261,414262,426023,426397,426398,427377,431947,440260,454453,454454,454455,454456,454457,454458,455073,455263,455373,471560,473664,473665,473666,474198,474199,492109,492177,496621,496662,510213,540410,540459,542523,544156,544157,545155,630493,630494,670695},
	}
	Banks["ie_bankofireland"] = &Bank{
		"bankofireland","ie","Bank Of Ireland","Bank Of Ireland","https://www.bankofireland.com","#013c5c",
		[]int{401773,404744,404745,405926,405927,408550,408573,420759,425717,427337,431938,431940,431941,431944,448449,453925,454679,455032,459922,472710,476210,476211,477575,477576,477577,477578,477579,477580,477581,477582,477583,477584,477585,477586,477587,477588,491086,493413,494050,494051,498833,510110,510233,520332,520333,540917,542509,542526,543267,543270,545645,546258,549897,554545,560390},
	}
	Banks["ir_ansar"] = &Bank{
		"ansar","ir","بانک انصار","Ansar Bank","http://ansarbank.com/","#b80710",
		[]int{627381},
	}
	Banks["ir_central"] = &Bank{
		"central","ir","بانک مرکزی جمهوری اسلامی ایران","Central Bank of the Islamic Republic of Iran","http://www.cbi.ir/","#22304d",
		[]int{636795},
	}
	Banks["ir_day"] = &Bank{
		"day","ir","بانک دی","Day Bank","http://bank-day.org/","#008b9a",
		[]int{502938},
	}
	Banks["ir_eghtesadnovin"] = &Bank{
		"eghtesadnovin","ir","بانک اقتصاد نوین","Eghtesad Novin Bank","http://www.enbank.ir/","#b42728",
		[]int{627412,627884,639194},
	}
	Banks["ir_exportdevelopment"] = &Bank{
		"exportdevelopment","ir","بانک توسعه صادرات","Export Development Bank","https://www.edbi.ir/","#006000",
		[]int{207177,627648},
	}
	Banks["ir_ghavamin"] = &Bank{
		"ghavamin","ir","بانک قوامین","Ghavamin Bank","http://www.ghbi.ir/","#00a652",
		[]int{639599},
	}
	Banks["ir_hekmatiranian"] = &Bank{
		"hekmatiranian","ir","بانک حکمت ایرانیان","Hekmat Iranian Bank","http://www.hibank24.ir/","#1d538d",
		[]int{636949},
	}
	Banks["ir_industryandmine"] = &Bank{
		"industryandmine","ir","بانک صنعت و معدن","Bank of Industry and Mine","http://www.bim.ir/","#e7c07d",
		[]int{627961},
	}
	Banks["ir_iranzamin"] = &Bank{
		"iranzamin","ir","بانک ایران زمین","Iran Zamin Bank","http://www.izbank.ir/","#4c2a94",
		[]int{505785},
	}
	Banks["ir_karafarin"] = &Bank{
		"karafarin","ir","بانک کارآفرین","Karafarin Bank","http://www.karafarinbank.ir/","#1f764c",
		[]int{502910,627488},
	}
	Banks["ir_keshavarzi"] = &Bank{
		"keshavarzi","ir","بانک کشاورزی","Keshavarzi Bank","http://www.bki.ir/","#005902",
		[]int{603770,639217},
	}
	Banks["ir_maskan"] = &Bank{
		"maskan","ir","بانک مسکن","Maskan Bank","http://bank-maskan.ir/","#e84b0a",
		[]int{628023},
	}
	Banks["ir_mellat"] = &Bank{
		"mellat","ir","بانک ملت","Mellat Bank","http://www.bankmellat.ir/","#d81736",
		[]int{610433,991975},
	}
	Banks["ir_melli"] = &Bank{
		"melli","ir","بانک ملی ایران","Melli Bank","http://bmi.ir/","#dc3c00",
		[]int{603799},
	}
	Banks["ir_parsian"] = &Bank{
		"parsian","ir","بانک پارسیان","Parsian Bank","http://www.parsian-bank.ir/","#ffc41e",
		[]int{622106},
	}
	Banks["ir_pasargad"] = &Bank{
		"pasargad","ir","بانک پاسارگاد","Pasargad Bank","http://bpi.ir/","#ffc41e",
		[]int{502229,639347},
	}
	Banks["ir_post"] = &Bank{
		"post","ir","پست بانک","Post Bank","https://ib.postbank.ir/","#006600",
		[]int{627760},
	}
	Banks["ir_qarzolqasanehmehriran"] = &Bank{
		"qarzolqasanehmehriran","ir","بانک قرض\u200cالحسنه مهر ایرانیان","Qarzol-Hasaneh Mehr Iran Bank","http://www.qmb.ir/","#009738",
		[]int{606373},
	}
	Banks["ir_refah"] = &Bank{
		"refah","ir","بانک رفاه","Refah Bank","http://www.rb24.ir/","#004b84",
		[]int{589463},
	}
	Banks["ir_saderat"] = &Bank{
		"saderat","ir","بانک صادرات","Saderat Bank","http://bsi.ir/","#005da2",
		[]int{603769},
	}
	Banks["ir_saman"] = &Bank{
		"saman","ir","بانک سامان","Saman Bank","https://www.sb24.com/","#009ddc",
		[]int{621986},
	}
	Banks["ir_sarmayeh"] = &Bank{
		"sarmayeh","ir","بانک سرمایه","Sarmayeh Bank","http://www.sbank.ir/","#000064",
		[]int{639607},
	}
	Banks["ir_sepah"] = &Bank{
		"sepah","ir","بانک سپه","Sepah Bank","https://www.ebanksepah.ir/","#e67a00",
		[]int{589210},
	}
	Banks["ir_shahr"] = &Bank{
		"shahr","ir","بانک شهر","Shahr Bank","http://shahr-bank.ir/","#ed1b24",
		[]int{502806,504706},
	}
	Banks["ir_sina"] = &Bank{
		"sina","ir","بانک سینا","Sina Bank","http://www.sinabank.ir/","#07519a",
		[]int{639346},
	}
	Banks["ir_tejarat"] = &Bank{
		"tejarat","ir","بانک تجارت","Tejarat Bank","http://www.tejaratbank.ir/","#2f4a98",
		[]int{585983,627353},
	}
	Banks["ir_toseetaavon"] = &Bank{
		"toseetaavon","ir","بانک توسعه تعاون","Tose'e Ta'avon Bank","https://www.ttbank.ir/","#50b5d1",
		[]int{502908},
	}
	Banks["ir_tourism"] = &Bank{
		"tourism","ir","بانک گردشگری","Tourism Bank","http://www.tourismbank.ir/","#b11116",
		[]int{505416},
	}
	Banks["kg_demirbank"] = &Bank{
		"demirbank","kg","Демир Кыргыз Интернэшнл Банк","DemirBank","http://demirbank.kg/","#c21f1a",
		[]int{417219,417221},
	}
	Banks["kz_bcc"] = &Bank{
		"bcc","kz","Банк ЦентрКредит","Bank CenterCredit","https://www.bcc.kz/","#008047",
		[]int{466720},
	}
	Banks["kz_kazkom"] = &Bank{
		"kazkom","kz","Казкоммерцбанк","Kazkommertsbank","http://kkb.kz/","#008fb4",
		[]int{400303,440564,557834},
	}
	Banks["kz_sberbank"] = &Bank{
		"sberbank","kz","Сбербанк","Sberbank","https://www.sberbank.kz","#309c0b",
		[]int{426344},
	}
	Banks["nz_bnz"] = &Bank{
		"bnz","nz","Bank of New Zealand","Bank of New Zealand","https://www.bnz.co.nz/","#002f6b",
		[]int{428455},
	}
	Banks["pl_deutsche-bank-polska"] = &Bank{
		"deutsche-bank-polska","pl","Deutsche Bank Polska","Deutsche Bank Polska","https://www.deutschebank.pl/","#00abdf",
		[]int{401801,474168,474170,552044},
	}
	Banks["pl_ing-bsk"] = &Bank{
		"ing-bsk","pl","ING Bank Slaski","ING Bank Slaski","https://www.ingbank.pl/","#ff6200",
		[]int{416040,448379,472817,524382},
	}
	Banks["pl_millennium"] = &Bank{
		"millennium","pl","Bank Millennium","Bank Millennium","https://www.bankmillennium.pl/","#c82059",
		[]int{512748,512792,534714,542757},
	}
	Banks["pl_raiffeisen-polbank"] = &Bank{
		"raiffeisen-polbank","pl","Raiffeisen Polbank","Raiffeisen Polbank","http://raiffeisenpolbank.com/","#fff533",
		[]int{402637,429046,543310},
	}
	Banks["pt_bpi"] = &Bank{
		"bpi","pt","Banco Português de Investimento (BPI)","Portuguese Investment Bank","www.bancobpi.pt","#ff8d00",
		[]int{415159},
	}
	Banks["pt_millenniumbcp"] = &Bank{
		"millenniumbcp","pt","Millennium BCP","Millennium BCP","www.millenniumbcp.pt","#e7046b",
		[]int{676938},
	}
	Banks["rs_societegeneralebanka"] = &Bank{
		"societegeneralebanka","rs","Societe Generale Banka Srbija Ad Beograd","Societe Generale Banka Srbija Ad Beograd","http://www.societegenerale.rs/","#ed1a3a",
		[]int{425340},
	}
	Banks["ru_akbarsbank"] = &Bank{
		"akbarsbank","ru","Ак Барс Банк","Ak Bars Bank","https://www.akbars.ru/","#019d3c",
		[]int{424436,424438,424439,424440,424447,520935,520985,552958},
	}
	Banks["ru_alfabank"] = &Bank{
		"alfabank","ru","Альфа-Банк","Alfa-Bank","https://alfabank.ru/","#f22f17",
		[]int{408396,408397,415428,415429,415481,415482,426101,426113,428905,428906,45841,477964,479004,479087,510126,519778,521178,548601,548655,548673,548674,552175,552186,555933,555947,555949,555957,676371},
	}
	Banks["ru_avangard"] = &Bank{
		"avangard","ru","Авангард","Avangard","https://avangard.ru/","#187543",
		[]int{522224},
	}
	Banks["ru_citibank"] = &Bank{
		"citibank","ru","Ситибанк","Citibank","https://www.citibank.ru/","#0088cf",
		[]int{419349,427760,427761,520306,527594},
	}
	Banks["ru_gazprombank"] = &Bank{
		"gazprombank","ru","Газпромбанк","Gazprombank","http://gazprombank.ru/","#00448b",
		[]int{487415,677585},
	}
	Banks["ru_mdm"] = &Bank{
		"mdm","ru","МДМ Банк","MDM Bank","https://www.mdm.ru/","#820729",
		[]int{554372},
	}
	Banks["ru_promsvyazbank"] = &Bank{
		"promsvyazbank","ru","Промсвязьбанк","Promsvyazbank","http://www.psbank.ru/","#274193",
		[]int{411791,447818,472345,520373,554759,554781},
	}
	Banks["ru_qiwi"] = &Bank{
		"qiwi","ru","Киви","QIWI","https://qiwi.ru","#ff9f00",
		[]int{469395,489049},
	}
	Banks["ru_raiffeisen"] = &Bank{
		"raiffeisen","ru","Райффайзенбанк","Raiffeisenbank","http://www.raiffeisen.ru/","#fff300",
		[]int{446917,447624,462729,462730,462758,510069,510070,515876,533594},
	}
	Banks["ru_renaissance"] = &Bank{
		"renaissance","ru","Ренессанс Капитал","Renaissance Capital","http://www.rencap.ru/","#343434",
		[]int{520905},
	}
	Banks["ru_rocketbank"] = &Bank{
		"rocketbank","ru","Рокетбанк","Rocketbank","https://rocketbank.ru/","#fb8000",
		[]int{532130},
	}
	Banks["ru_rosbank"] = &Bank{
		"rosbank","ru","Росбанк","Rosbank","https://www.rosbank.ru/","#e60028",
		[]int{554761},
	}
	Banks["ru_rosevrobank"] = &Bank{
		"rosevrobank","ru","РосЕвроБанк","RosEvro Bank","http://www.rosevrobank.ru/","#6c2870",
		[]int{418907},
	}
	Banks["ru_rsb"] = &Bank{
		"rsb","ru","Банк Русский Стандарт","Russian Standard Bank","https://www.rsb.ru/","#930d2d",
		[]int{510047,513691,542048},
	}
	Banks["ru_sberbank"] = &Bank{
		"sberbank","ru","Сбербанк","Sberbank","https://www.sberbank.ru","#309c0b",
		[]int{427466,427467,427468,427469,427470,427472,427475,427477,427601,427602,427606,427609,427610,427616,427620,427622,427630,427631,427637,427638,427640,427642,427644,427648,42765,427660,427663,427664,427667,427668,427669,427672,427680,427683,427684,427685,427686,427687,427688,427901,427916,427938,427940,427955,533669,54690,546910,546916,546917,546918,54693,546940,546942,546944,546945,546949,54695,546960,546962,546967,546969,546977,546999,547938,548401,548406,548438,639002,676196,67758},
	}
	Banks["ru_siab"] = &Bank{
		"siab","ru","СИАБ","SIAB","https://siab.ru/","#0f255c",
		[]int{531544},
	}
	Banks["ru_sovcombank"] = &Bank{
		"sovcombank","ru","Совкомбанк","Sovcombank","https://sovcombank.ru/","#134c7c",
		[]int{533595},
	}
	Banks["ru_tinkoff"] = &Bank{
		"tinkoff","ru","Тинькофф Банк","Tinkoff Bank","https://www.tinkoff.ru/","#ffdd2d",
		[]int{437772,437773,521324,548387,553691},
	}
	Banks["ru_trust"] = &Bank{
		"trust","ru","Траст","Trust","http://www.trust.ru/","#f8ab21",
		[]int{456515,554521},
	}
	Banks["ru_uniastrumbank"] = &Bank{
		"uniastrumbank","ru","Юниаструм Банк","Uniastrum Bank","https://www.uniastrum.ru","#00718f",
		[]int{403330},
	}
	Banks["ru_unicreditbank"] = &Bank{
		"unicreditbank","ru","ЮниКредит Банк","UniCredit Bank","https://www.unicreditbank.ru/","#e3001a",
		[]int{485608,489042,490855,531344},
	}
	Banks["ru_uralsibbank"] = &Bank{
		"uralsibbank","ru","УралСиб Банк","UralSib Bank","https://www.bankuralsib.ru/","#53b548",
		[]int{424204,440666},
	}
	Banks["ru_vozrozhdenie"] = &Bank{
		"vozrozhdenie","ru","Возрождение","Vozrozhdenie","http://www.vbank.ru/","#224f84",
		[]int{526469},
	}
	Banks["ru_vpb"] = &Bank{
		"vpb","ru","ВПБ","VPB","https://www.vpb.ru","#0071b9",
		[]int{534134},
	}
	Banks["ru_vtb24"] = &Bank{
		"vtb24","ru","ВТБ 24","VTB 24","http://www.vtb24.ru","#00498f",
		[]int{418869,427229,447520,46223,465206,471487,522598,527883,535082,542033,554386,554393},
	}
	Banks["ru_yandex"] = &Bank{
		"yandex","ru","Яндекс.Деньги","Yandex Money","https://money.yandex.ru/","#000000",
		[]int{510621,518901},
	}
	Banks["tr_akbank"] = &Bank{
		"akbank","tr","Akbank","Akbank","https://akbank.com","#4285f4",
		[]int{413252,425669,432072,435508,435509,493837,512754,520932,521807,524347,542110,552608,552609,553056,557113,557829},
	}
	Banks["tr_albaraka"] = &Bank{
		"albaraka","tr","al Baraka","al Baraka","https://albaraka.com.tr/","#4285f4",
		[]int{417715,432284,432285,534264,547234,548232},
	}
	Banks["tr_bankasya"] = &Bank{
		"bankasya","tr","Bank Asya","Bank Asya","https://bankasya.com.tr/","#4285f4",
		[]int{402275,402276,402280,416987,441033,515849,524384,527585,529462,531334,547799,552529},
	}
	Banks["tr_denizbank"] = &Bank{
		"denizbank","tr","Denizbank","Denizbank","http://denizbank.com/","#4285f4",
		[]int{403134,408625,409070,411924,423667,424360,424361,426391,426392,441139,450050,450051,460345,460347,472914,489456,510063,510118,510119,512017,512117,514924,520019,520303,521376,543358,543400,543427,544127,544445,544460,546764,547712,549220,554483,558514},
	}
	Banks["tr_finansbank"] = &Bank{
		"finansbank","tr","Finansbank","Finansbank","https://finansbank.com.tr/","#4285f4",
		[]int{402277,402278,402563,403082,409364,410147,413583,414388,415565,422376,423277,423398,427311,435653,441007,444029,499850,499851,499852,519324,521022,521836,529572,531157,545120,545616,545847,547567,547800},
	}
	Banks["tr_garanti"] = &Bank{
		"garanti","tr","Garanti Bankasi","Garanti Bankasi","http://garanti.com.tr/","#4285f4",
		[]int{403280,403666,404308,413836,426886,426887,426888,427314,427315,428220,428221,432154,448472,461668,462274,467293,467294,467295,474151,482489,482490,482491,486567,487074,487075,489478,490175,492186,492187,492193,493845,514915,520097,520922,520940,520988,521368,521824,521825,521832,522204,528939,528956,533169,534261,540036,540037,540226,540227,540669,540709,541865,542030,544078,545102,546001,547302,552095,553130,554796,554960,557023,557945,558699},
	}
	Banks["tr_halkbank"] = &Bank{
		"halkbank","tr","Halk Bankası","Halk Bankası","https://halkbank.com.tr/","#4285f4",
		[]int{415514,492094,492095,498852,510056,521378,540435,543081,552879},
	}
	Banks["tr_hsbc"] = &Bank{
		"hsbc","tr","HSBC","HSBC","https://hsbc.com.tr/","#4285f4",
		[]int{405913,405917,405918,409071,422629,424909,428240,496019,510005,512651,519399,521045,522054,525413,525795,540643,542254,545183,550472,550473,552143,556030,556031,556033,556034,556665},
	}
	Banks["tr_ingbank"] = &Bank{
		"ingbank","tr","ING Bank","ING Bank","https://ingbank.com.tr/","#4285f4",
		[]int{400684,408579,414070,420322,420323,420324,455571,480296,490805,490806,490807,510151,532443,540024,540025,542029,542605,542965,542967,547765,548819,554297,554570},
	}
	Banks["tr_isbank"] = &Bank{
		"isbank","tr","Türkiye İş Bankası","Türkiye İş Bankası","http://isbank.com.tr/","#4285f4",
		[]int{418342,418343,418344,418345,450803,454318,454358,454359,454360,510152,540667,540668,543771,552096,553058},
	}
	Banks["tr_kuveytturk"] = &Bank{
		"kuveytturk","tr","Kuveyt Türk","Kuveyt Türk","https://kuveytturk.com.tr/","#4285f4",
		[]int{402589,402590,402592,403360,403810,410555,410556,424487,431024,511660,512595,518896,520180,525312,547564},
	}
	Banks["tr_sekerbank"] = &Bank{
		"sekerbank","tr","Şeker Bank","Şeker Bank","http://sekerbank.com.tr/","#4285f4",
		[]int{403836,411156,411157,411158,411159,411160,433383,433384,494063,494064,521394,521827,525404,530866,539703,547311,549208,549394},
	}
	Banks["tr_teb"] = &Bank{
		"teb","tr","Türk Ekonomi Bankası","Türk Ekonomi Bankası","http://teb.com.tr/","#4285f4",
		[]int{402458,402459,406015,427707,440247,440273,440293,440294,479227,489494,489495,489496,510138,510139,510221,512753,512803,524346,524839,524840,528920,530853,545124,553090},
	}
	Banks["tr_turkiyefinans"] = &Bank{
		"turkiyefinans","tr","Türkiye Finans","Türkiye Finans","https://turkiyefinans.com.tr/","#4285f4",
		[]int{404952,411685,428462,435627,435628,521848,537719,549294},
	}
	Banks["tr_vakifbank"] = &Bank{
		"vakifbank","tr","Vakıf Bank","Vakıf Bank","https://vakifbank.com.tr/","#4285f4",
		[]int{402940,409084,411724,411942,411943,411944,411979,415792,416757,428945,493840,493841,493846,520017,540045,540046,542119,542798,542804,547244,552101},
	}
	Banks["tr_yapikredi"] = &Bank{
		"yapikredi","tr","Yapı ve Kredi Bankası","Yapı ve Kredi Bankası","http://yapikredi.com.tr/","#4285f4",
		[]int{404809,446212,450634,455359,477959,479794,479795,491205,491206,492128,492130,492131,510054,540061,540062,540063,540122,540129,542117,545103,552645,552659,554422},
	}
	Banks["tr_ziraatbankasi"] = &Bank{
		"ziraatbankasi","tr","Ziraat Bankası","Ziraat Bankası","http://ziraatbank.com.tr","#4285f4",
		[]int{413226,444676,444677,444678,453955,453956,454671,454672,454673,454674,454894,540130,540134,541001,541033,542374,547287},
	}
	Banks["tw_anz"] = &Bank{
		"anz","tw","澳盛(台灣)商業銀行","ANZ Bank (Taiwan)","http://www.anz.tw","#003f67",
		[]int{442355,468236,515687,539657,552129},
	}
	Banks["tw_cathaybk"] = &Bank{
		"cathaybk","tw","國泰世華商業銀行","Cathay United Bank","https://www.cathaybk.com.tw","#11a847",
		[]int{402310,428368,428430,474832,517478,519936,524106,543375,546696,552197,552579},
	}
	Banks["tw_citibank"] = &Bank{
		"citibank","tw","花旗(台灣)銀行","Citibank Taiwan","https://www.citibank.com.tw","#0088cf",
		[]int{428440,431178,456313,456318,463670,543374,552000},
	}
	Banks["tw_ctbcbank"] = &Bank{
		"ctbcbank","tw","中國信託商業銀行","CTBC Bank","https://www.ctbcbank.com","#007d7d",
		[]int{400361,406586,416217,418230,421013,421016,421599,430451,431195,447757,451873,456301,456319,481537,483515,515352,517357,523953,533785,533786,537675,537805,540829,543372,546697,547785,552049,558888,559569},
	}
	Banks["tw_esunbank"] = &Bank{
		"esunbank","tw","玉山銀行","E.SUN Bank","https://www.esunbank.com.tw","#009d99",
		[]int{404381,409679,416210,439237,457966,457967,461160,464961,483138,486606,511098,512211,512308,515717,515879,515898,519933,519934,519935,520063,520119,521129,522225,522831,523903,523976,524255,528602,528950,539228,539231,542462,542700,552199,555029},
	}
	Banks["tw_fubon"] = &Bank{
		"fubon","tw","台北富邦銀行","Taipei Fubon Bank","https://www.fubon.com","#007c97",
		[]int{400357,402987,410523,436314,451445,457952,457953,457956,457957,462410,463781,469531,472634,472635,481509,492496,498713,515709,515716,517951,523764,524108,524132,524264,540945,540966,543377,543381,546600,552046,552345,555033,558805},
	}
	Banks["tw_hncb"] = &Bank{
		"hncb","tw","華南銀行","Hua Nan Bank","http://www.hncb.com.tw","#ce1624",
		[]int{486375},
	}
	Banks["tw_scsb"] = &Bank{
		"scsb","tw","上海商業儲蓄銀行","he Shanghai Commercial & Savings Bank, Ltd. ","http://www.scsb.com.tw","#c42607",
		[]int{403879,456314,456317,514738,540853},
	}
	Banks["tw_tcbbank"] = &Bank{
		"tcbbank","tw","台中商業銀行","Taichung Bank","https://www.tcbbank.com.tw","#ce1624",
		[]int{558888},
	}
	Banks["ua_alfabank"] = &Bank{
		"alfabank","ua","Альфа-Банк","Alfa-Bank","https://alfabank.ua/","#ed1c2a",
		[]int{410247,422605,423719,536364,543259,544904,547994},
	}
	Banks["ua_credit-agricole"] = &Bank{
		"credit-agricole","ua","Креді Агріколь Банк (Україна)","Credit Agricole (Ukraine)","https://credit-agricole.ua/","#258b92",
		[]int{411948,512439,519745,520662,525822},
	}
	Banks["ua_diamant"] = &Bank{
		"diamant","ua","Діамантбанк","Diamantbank","http://diamantbank.ua/","#0a2855",
		[]int{423174,429453,467936},
	}
	Banks["ua_kredobank"] = &Bank{
		"kredobank","ua","Кредобанк","KredoBank","http://www.kredobank.com.ua/","#275797",
		[]int{438253},
	}
	Banks["ua_oschadbank"] = &Bank{
		"oschadbank","ua","Ощадбанк","Oschadbank","http://www.oschadbank.ua/","#06503c",
		[]int{515662,545265,555952,557744},
	}
	Banks["ua_otp"] = &Bank{
		"otp","ua","OTP Bank","OTP Bank","https://www.otpbank.com.ua/","#a3d044",
		[]int{406760,406788,411749,416955,436323,471362,528534,533194,546886,547266},
	}
	Banks["ua_pravex"] = &Bank{
		"pravex","ua","ПРАВЕКС-БАНК","PRAVEX-BANK","https://www.pravex.com.ua/","#701616",
		[]int{404764,405247,409089,409388,409625,409881,461278,536444,541701,544142,552670,676274},
	}
	Banks["ua_privatbank"] = &Bank{
		"privatbank","ua","ПриватБанк","Privatbank","https://privatbank.ua/","#498d2f",
		[]int{410653,414949,414960,414962,414963,423396,424600,432575,434158,440509,440535,440588,458122,462708,473118,473121,475538,476339,516874,516875,516933,521153,521857,530217,532957,536354,544457,545709,552324,558335,558424},
	}
	Banks["ua_procredit"] = &Bank{
		"procredit","ua","ПроКредит Банк (Украина)","ProCredit Bank (Ukraine)","http://procreditbank.com.ua/","#da2027",
		[]int{417679,433424},
	}
	Banks["ua_pumb"] = &Bank{
		"pumb","ua","ПУМБ (Перший Український Міжнародний Банк)","FUIB (First Ukrainian International Bank)","https://pumb.ua","#d13239",
		[]int{404170,406660,412717,419316,428337,431437,482415,510117,510536,512647,516059,516750,518778,518985,520173,521327,521514,521762,521832,523778,524729,524814,525410,525493,526858,530186,530359,533686,533731,537003,538814,539715,540892,540953,541271,542171,542521,544044,544187,545219,547650,548343,548919,549385,550250,552154,552611,553045,553053,557104,558262,558381,558591,670604,670832,676622,676652,676700,677210,677567,677701},
	}
	Banks["ua_raiffeisen-aval"] = &Bank{
		"raiffeisen-aval","ua","Райффайзен Банк Аваль","Raiffeisen Bank Aval","https://www.aval.ua/","#fffb0b",
		[]int{413062,414953,418844,423843,462775,512070,512557,512572,516913,518808,529919,532457,535160,536502,542465,543155,543759,545253,548658,548968,552657,557723,557881},
	}
	Banks["ua_ukreximbank"] = &Bank{
		"ukreximbank","ua","Укрэксимбанк","Ukreximbank","https://www.eximb.com/","#0067a6",
		[]int{403520,406337,423297,427705,440272,478766,510140,515855,516882,518720,522473,522766,522872,523473,524876,527598,528953,529801,530004,530094,536270,539897,540381,540682,541740,541940,542031,543205,544510,547390,547734,547762,548390,548766,550617,552247,552571,552624,552735,553465,554966,557833,558311,558797,676895,676965,677046,677279,677325,677613},
	}
	Banks["ua_ukrgasbank"] = &Bank{
		"ukrgasbank","ua","УКРГАЗБАНК","Ukrgasbank","http://www.ukrgasbank.com/","#005596",
		[]int{401800,403471,510449,512100,512743,516854,517725,517818,518654,521833,522654,522797,522967,524872,525672,525792,527412,528812,531007,531725,532343,539845,540305,540617,540637,543438,547313,547914,548883,549146,549582,549824,549896,553391,557110,670651,670669,670877,671103,677237,677490},
	}
	Banks["ua_ukrsibbank"] = &Bank{
		"ukrsibbank","ua","УКРСИББАНК","Ukrsibbank","https://ukrsibbank.com/","#04703d",
		[]int{407361,407368,417232,427830,434343,516930,529578,533037,535129},
	}
	Banks["ua_ukrsotsbank"] = &Bank{
		"ukrsotsbank","ua","Укрсоцбанк","Ukrsotsbank","https://ukrsotsbank.com/","#a50931",
		[]int{402191,417284,431153,476009,487412},
	}
	Banks["ua_unicreditbank"] = &Bank{
		"unicreditbank","ua","ЮниКредит Банк","UniCredit Bank","https://unicredit.ua/","#e3001a",
		[]int{408339,430333,430364,432109,487410,510467,512789,515794,518282,525759,546644,553091},
	}
	Banks["us_ally"] = &Bank{
		"ally","us","Ally","Ally","https://www.ally.com","#650360",
		[]int{557552},
	}
	Banks["us_americanexpress"] = &Bank{
		"americanexpress","us","American Express","American Express","https://www.americanexpress.com","#004fc5",
		[]int{370370,370371,370372,370373,370374,370375,370376,370377,370378,370379,370700,370701,370702,370703,370704,370705,370706,370707,370708,370709,370710,370711,370712,370713,370714,370715,370716,370717,370718,370719,370720,370721,370722,370723,370724,370725,370726,370727,370728,370729,370730,370731,370732,370733,370734,370735,370736,370737,370738,370739,370740,370741,370742,370743,370744,370745,370746,370747,370748,370749,370750,370751,370752,370753,370754,370755,370756,370757,370758,370759,370760,370761,370762,370763,370764,370765,370766,370767,370768,370769,370770,370771,370772,370773,370774,370775,370776,370777,370778,370779,370780,370781,370782,370783,370784,370785,370786,370787,370788,370789,370790,370791,370792,370793,370794,370795,370796,370797,370798,370799,371000,371001,371002,371003,371004,371005,371006,371007,371008,371009,371010,371011,371012,371013,371014,371015,371016,371017,371018,371019,371020,371021,371022,371023,371024,371025,371026,371027,371028,371029,371030,371031,371032,371033,371034,371035,371036,371037,371038,371039,371040,371041,371042,371043,371044,371045,371046,371047,371048,371049,371050,371051,371052,371053,371054,371055,371056,371057,371058,371059,371060,371061,371062,371063,371064,371065,371066,371067,371068,371069,371070,371071,371072,371073,371074,371075,371076,371077,371078,371079,371080,371081,371082,371083,371084,371085,371086,371087,371088,371089,371090,371091,371092,371093,371094,371095,371096,371097,371098,371099,371100,371101,371102,371103,371104,371105,371106,371107,371108,371109,371110,371111,371112,371113,371114,371115,371116,371117,371118,371119,371120,371121,371122,371123,371124,371125,371126,371127,371128,371129,371130,371131,371132,371133,371134,371135,371136,371137,371138,371139,371140,371141,371142,371143,371144,371145,371146,371147,371148,371149,371150,371151,371152,371153,371154,371155,371156,371157,371158,371159,371160,371161,371162,371163,371164,371165,371166,371167,371168,371169,371170,371171,371172,371173,371174,371175,371176,371177,371178,371179,371180,371181,371182,371183,371184,371185,371186,371187,371188,371189,371190,371191,371192,371193,371194,371195,371196,371197,371198,371199,371200,371201,371202,371203,371204,371205,371206,371207,371208,371209,371210,371211,371212,371213,371214,371215,371216,371217,371218,371219,371220,371221,371222,371223,371224,371225,371226,371227,371228,371229,371230,371231,371232,371233,371234,371235,371236,371237,371238,371239,371240,371241,371242,371243,371244,371245,371246,371247,371248,371249,371250,371251,371252,371253,371254,371255,371256,371257,371258,371259,371260,371261,371262,371263,371264,371265,371266,371267,371268,371269,371270,371271,371272,371273,371274,371275,371276,371277,371278,371279,371280,371281,371282,371283,371284,371285,371286,371287,371288,371289,371290,371291,371292,371293,371294,371295,371296,371297,371298,371299,371300,371301,371302,371303,371304,371305,371306,371307,371308,371309,371310,371311,371312,371313,371314,371315,371316,371317,371318,371319,371320,371321,371322,371323,371324,371325,371326,371327,371328,371329,371330,371331,371332,371333,371334,371335,371336,371337,371338,371339,371340,371341,371342,371343,371344,371345,371346,371347,371348,371349,371350,371351,371352,371353,371354,371355,371356,371357,371358,371359,371360,371361,371362,371363,371364,371365,371366,371367,371368,371369,371370,371371,371372,371373,371374,371375,371376,371377,371378,371379,371380,371381,371382,371383,371384,371385,371386,371387,371388,371389,371390,371391,371392,371393,371394,371395,371396,371397,371398,371399,371400,371401,371402,371403,371404,371405,371406,371407,371408,371409,371410,371411,371412,371413,371414,371415,371416,371417,371418,371419,371420,371421,371422,371423,371424,371425,371426,371427,371428,371429,371430,371431,371432,371433,371434,371435,371436,371437,371438,371439,371440,371441,371442,371443,371444,371445,371446,371447,371448,371449,371450,371451,371452,371453,371454,371455,371456,371457,371458,371459,371460,371461,371462,371463,371464,371465,371466,371467,371468,371469,371470,371471,371472,371473,371474,371475,371476,371477,371478,371479,371480,371481,371482,371483,371484,371485,371486,371487,371488,371489,371490,371491,371492,371493,371494,371495,371496,371497,371498,371499,371500,371501,371502,371503,371504,371505,371506,371507,371508,371509,371510,371511,371512,371513,371514,371515,371516,371517,371518,371519,371520,371521,371522,371523,371524,371525,371526,371527,371528,371529,371530,371531,371532,371533,371534,371535,371536,371537,371538,371539,371540,371541,371542,371543,371544,371545,371546,371547,371548,371549,371550,371551,371552,371553,371554,371555,371556,371557,371558,371559,371560,371561,371562,371563,371564,371565,371566,371567,371568,371569,371570,371571,371572,371573,371574,371575,371576,371577,371578,371579,371580,371581,371582,371583,371584,371585,371586,371587,371588,371589,371590,371591,371592,371593,371594,371595,371596,371597,371598,371599,371600,371601,371602,371603,371604,371605,371606,371607,371608,371609,371610,371611,371612,371613,371614,371615,371616,371617,371618,371619,371620,371621,371622,371623,371624,371625,371626,371627,371628,371629,371630,371631,371632,371633,371634,371635,371636,371637,371638,371639,371640,371641,371642,371643,371644,371645,371646,371647,371648,371649,371650,371651,371652,371653,371654,371655,371656,371657,371658,371659,371660,371661,371662,371663,371664,371665,371666,371667,371668,371669,371670,371671,371672,371673,371674,371675,371676,371677,371678,371679,371680,371681,371682,371683,371684,371685,371686,371687,371688,371689,371690,371691,371692,371693,371694,371695,371696,371697,371698,371699,371700,371701,371702,371703,371704,371705,371706,371707,371708,371709,371710,371711,371712,371713,371714,371715,371716,371717,371718,371719,371720,371721,371722,371723,371724,371725,371726,371727,371728,371729,371730,371731,371732,371733,371734,371735,371736,371737,371738,371739,371740,371741,371742,371743,371744,371745,371746,371747,371748,371749,371750,371751,371752,371753,371754,371755,371756,371757,371758,371759,371760,371761,371762,371763,371764,371765,371766,371767,371768,371769,371770,371771,371772,371773,371774,371775,371776,371777,371778,371779,371780,371781,371782,371783,371784,371785,371786,371787,371788,371789,371790,371791,371792,371793,371794,371795,371796,371797,371798,371799,371800,371801,371802,371803,371804,371805,371806,371807,371808,371809,371810,371811,371812,371813,371814,371815,371816,371817,371818,371819,371820,371821,371822,371823,371824,371825,371826,371827,371828,371829,371830,371831,371832,371833,371834,371835,371836,371837,371838,371839,371840,371841,371842,371843,371844,371845,371846,371847,371848,371849,371850,371851,371852,371853,371854,371855,371856,371857,371858,371859,371860,371861,371862,371863,371864,371865,371866,371867,371868,371869,371870,371871,371872,371873,371874,371875,371876,371877,371878,371879,371880,371881,371882,371883,371884,371885,371886,371887,371888,371889,371890,371891,371892,371893,371894,371895,371896,371897,371898,371899,371900,371901,371902,371903,371904,371905,371906,371907,371908,371909,371910,371911,371912,371913,371914,371915,371916,371917,371918,371919,371920,371921,371922,371923,371924,371925,371926,371927,371928,371929,371930,371931,371932,371933,371934,371935,371936,371937,371938,371939,371940,371941,371942,371943,371944,371945,371946,371947,371948,371949,371950,371951,371952,371953,371954,371955,371956,371957,371958,371959,371960,371961,371962,371963,371964,371965,371966,371967,371968,371969,371970,371971,371972,371973,371974,371975,371976,371977,371978,371979,371980,371981,371982,371983,371984,371985,371986,371987,371988,371989,371990,371991,371992,371993,371994,371995,371996,371997,371998,371999,372000,372001,372002,372003,372004,372005,372006,372007,372008,372009,372010,372011,372012,372013,372014,372015,372016,372017,372018,372019,372020,372021,372022,372023,372024,372025,372026,372027,372028,372029,372030,372031,372032,372033,372034,372035,372036,372037,372038,372039,372040,372041,372042,372043,372044,372045,372046,372047,372048,372049,372050,372051,372052,372053,372054,372055,372056,372057,372058,372059,372060,372061,372062,372063,372064,372065,372066,372067,372068,372069,372070,372071,372072,372073,372074,372075,372076,372077,372078,372079,372080,372081,372082,372083,372084,372085,372086,372087,372088,372089,372090,372091,372092,372093,372094,372095,372096,372097,372098,372099,372100,372101,372102,372103,372104,372105,372106,372107,372108,372109,372110,372111,372112,372113,372114,372115,372116,372117,372118,372119,372120,372121,372122,372123,372124,372125,372126,372127,372128,372129,372130,372131,372132,372133,372134,372135,372136,372137,372138,372139,372140,372141,372142,372143,372144,372145,372146,372147,372148,372149,372150,372151,372152,372153,372154,372155,372156,372157,372158,372159,372160,372161,372162,372163,372164,372165,372166,372167,372168,372169,372170,372171,372172,372173,372174,372175,372176,372177,372178,372179,372180,372181,372182,372183,372184,372185,372186,372187,372188,372189,372190,372191,372192,372193,372194,372195,372196,372197,372198,372199,372200,372201,372202,372203,372204,372205,372206,372207,372208,372209,372210,372211,372212,372213,372214,372215,372216,372217,372218,372219,372220,372221,372222,372223,372224,372225,372226,372227,372228,372229,372230,372231,372232,372233,372234,372235,372236,372237,372238,372239,372240,372241,372242,372243,372244,372245,372246,372247,372248,372249,372250,372251,372252,372253,372254,372255,372256,372257,372258,372259,372260,372261,372262,372263,372264,372265,372266,372267,372268,372269,372270,372271,372272,372273,372274,372275,372276,372277,372278,372279,372280,372281,372282,372283,372284,372285,372286,372287,372288,372289,372290,372291,372292,372293,372294,372295,372296,372297,372298,372299,372300,372301,372302,372303,372304,372305,372306,372307,372308,372309,372310,372311,372312,372313,372314,372315,372316,372317,372318,372319,372320,372321,372322,372323,372324,372325,372326,372327,372328,372329,372330,372331,372332,372333,372334,372335,372336,372337,372338,372339,372340,372341,372342,372343,372344,372345,372346,372347,372348,372349,372350,372351,372352,372353,372354,372355,372356,372357,372358,372359,372360,372361,372362,372363,372364,372365,372366,372367,372368,372369,372370,372371,372372,372373,372374,372375,372376,372377,372378,372379,372380,372381,372382,372383,372384,372385,372386,372387,372388,372389,372390,372391,372392,372393,372394,372395,372396,372397,372398,372399,372400,372401,372402,372403,372404,372405,372406,372407,372408,372409,372410,372411,372412,372413,372414,372415,372416,372417,372418,372419,372420,372421,372422,372423,372424,372425,372426,372427,372428,372429,372430,372431,372432,372433,372434,372435,372436,372437,372438,372439,372440,372441,372442,372443,372444,372445,372446,372447,372448,372449,372450,372451,372452,372453,372454,372455,372456,372457,372458,372459,372460,372461,372462,372463,372464,372465,372466,372467,372468,372469,372470,372471,372472,372473,372474,372475,372476,372477,372478,372479,372480,372481,372482,372483,372484,372485,372486,372487,372488,372489,372490,372491,372492,372493,372494,372495,372496,372497,372498,372499,372500,372501,372502,372503,372504,372505,372506,372507,372508,372509,372510,372511,372512,372513,372514,372515,372516,372517,372518,372519,372520,372521,372522,372523,372524,372525,372526,372527,372528,372529,372530,372531,372532,372533,372534,372535,372536,372537,372538,372539,372540,372541,372542,372543,372544,372545,372546,372547,372548,372549,372550,372551,372552,372553,372554,372555,372556,372557,372558,372559,372560,372561,372562,372563,372564,372565,372566,372567,372568,372569,372570,372571,372572,372573,372574,372575,372576,372577,372578,372579,372580,372581,372582,372583,372584,372585,372586,372587,372588,372589,372590,372591,372592,372593,372594,372595,372596,372597,372598,372599,372600,372601,372602,372603,372604,372605,372606,372607,372608,372609,372610,372611,372612,372613,372614,372615,372616,372617,372618,372619,372620,372621,372622,372623,372624,372625,372626,372627,372628,372629,372630,372631,372632,372633,372634,372635,372636,372637,372638,372639,372640,372641,372642,372643,372644,372645,372646,372647,372648,372649,372650,372651,372652,372653,372654,372655,372656,372657,372658,372659,372660,372661,372662,372663,372664,372665,372666,372667,372668,372669,372670,372671,372672,372673,372674,372675,372676,372677,372678,372679,372680,372681,372682,372683,372684,372685,372686,372687,372688,372689,372690,372691,372692,372693,372694,372695,372696,372697,372698,372699,372700,372701,372702,372703,372704,372705,372706,372707,372708,372709,372710,372711,372712,372713,372714,372715,372716,372717,372718,372719,372720,372721,372722,372723,372724,372725,372726,372727,372728,372729,372730,372731,372732,372733,372734,372735,372736,372737,372738,372739,372740,372741,372742,372743,372744,372745,372746,372747,372748,372749,372750,372751,372752,372753,372754,372755,372756,372757,372758,372759,372760,372761,372762,372763,372764,372765,372766,372767,372768,372769,372770,372771,372772,372773,372774,372775,372776,372777,372778,372779,372780,372781,372782,372783,372784,372785,372786,372787,372788,372789,372790,372791,372792,372793,372794,372795,372796,372797,372798,372799,372800,372801,372802,372803,372804,372805,372806,372807,372808,372809,372810,372811,372812,372813,372814,372815,372816,372817,372818,372819,372820,372821,372822,372823,372824,372825,372826,372827,372828,372829,372830,372831,372832,372833,372834,372835,372836,372837,372838,372839,372840,372841,372842,372843,372844,372845,372846,372847,372848,372849,372850,372851,372852,372853,372854,372855,372856,372857,372858,372859,372860,372861,372862,372863,372864,372865,372866,372867,372868,372869,372870,372871,372872,372873,372874,372875,372876,372877,372878,372879,372880,372881,372882,372883,372884,372885,372886,372887,372888,372889,372890,372891,372892,372893,372894,372895,372896,372897,372898,372899,372900,372901,372902,372903,372904,372905,372906,372907,372908,372909,372910,372911,372912,372913,372914,372915,372916,372917,372918,372919,372920,372921,372922,372923,372924,372925,372926,372927,372928,372929,372930,372931,372932,372933,372934,372935,372936,372937,372938,372939,372940,372941,372942,372943,372944,372945,372946,372947,372948,372949,372950,372951,372952,372953,372954,372955,372956,372957,372958,372959,372960,372961,372962,372963,372964,372965,372966,372967,372968,372969,372970,372971,372972,372973,372974,372975,372976,372977,372978,372979,372980,372981,372982,372983,372984,372985,372986,372987,372988,372989,372990,372991,372992,372993,372994,372995,372996,372997,372998,372999,373000,373001,373002,373003,373004,373005,373006,373007,373008,373009,373010,373011,373012,373013,373014,373015,373016,373017,373018,373019,373020,373021,373022,373023,373024,373025,373026,373027,373028,373029,373030,373031,373032,373033,373034,373035,373036,373037,373038,373039,373040,373041,373042,373043,373044,373045,373046,373047,373048,373049,373050,373051,373052,373053,373054,373055,373056,373057,373058,373059,373060,373061,373062,373063,373064,373065,373066,373067,373068,373069,373070,373071,373072,373073,373074,373075,373076,373077,373078,373079,373080,373081,373082,373083,373084,373085,373086,373087,373088,373089,373090,373091,373092,373093,373094,373095,373096,373097,373098,373099,373100,373101,373102,373103,373104,373105,373106,373107,373108,373109,373110,373111,373112,373113,373114,373115,373116,373117,373118,373119,373120,373121,373122,373123,373124,373125,373126,373127,373128,373129,373130,373131,373132,373133,373134,373135,373136,373137,373138,373139,373140,373141,373142,373143,373144,373145,373146,373147,373148,373149,373150,373151,373152,373153,373154,373155,373156,373157,373158,373159,373160,373161,373162,373163,373164,373165,373166,373167,373168,373169,373170,373171,373172,373173,373174,373175,373176,373177,373178,373179,373180,373181,373182,373183,373184,373185,373186,373187,373188,373189,373190,373191,373192,373193,373194,373195,373196,373197,373198,373199,373200,373201,373202,373203,373204,373205,373206,373207,373208,373209,373210,373211,373212,373213,373214,373215,373216,373217,373218,373219,373220,373221,373222,373223,373224,373225,373226,373227,373228,373229,373230,373231,373232,373233,373234,373235,373236,373237,373238,373239,373240,373241,373242,373243,373244,373245,373246,373247,373248,373249,373250,373251,373252,373253,373254,373255,373256,373257,373258,373259,373260,373261,373262,373263,373264,373265,373266,373267,373268,373269,373270,373271,373272,373273,373274,373275,373276,373277,373278,373279,373280,373281,373282,373283,373284,373285,373286,373287,373288,373289,373290,373291,373292,373293,373294,373295,373296,373297,373298,373299,373300,373301,373302,373303,373304,373305,373306,373307,373308,373309,373310,373311,373312,373313,373314,373315,373316,373317,373318,373319,373320,373321,373322,373323,373324,373325,373326,373327,373328,373329,373330,373331,373332,373333,373334,373335,373336,373337,373338,373339,373340,373341,373342,373343,373344,373345,373346,373347,373348,373349,373350,373351,373352,373353,373354,373355,373356,373357,373358,373359,373360,373361,373362,373363,373364,373365,373366,373367,373368,373369,373370,373371,373372,373373,373374,373375,373376,373377,373378,373379,373380,373381,373382,373383,373384,373385,373386,373387,373388,373389,373390,373391,373392,373393,373394,373395,373396,373397,373398,373399,373400,373401,373402,373403,373404,373405,373406,373407,373408,373409,373410,373411,373412,373413,373414,373415,373416,373417,373418,373419,373420,373421,373422,373423,373424,373425,373426,373427,373428,373429,373430,373431,373432,373433,373434,373435,373436,373437,373438,373439,373440,373441,373442,373443,373444,373445,373446,373447,373448,373449,373450,373451,373452,373453,373454,373455,373456,373457,373458,373459,373460,373461,373462,373463,373464,373465,373466,373467,373468,373469,373470,373471,373472,373473,373474,373475,373476,373477,373478,373479,373480,373481,373482,373483,373484,373485,373486,373487,373488,373489,373490,373491,373492,373493,373494,373495,373496,373497,373498,373499,373500,373501,373502,373503,373504,373505,373506,373507,373508,373509,373510,373511,373512,373513,373514,373515,373516,373517,373518,373519,373520,373521,373522,373523,373524,373525,373526,373527,373528,373529,373530,373531,373532,373533,373534,373535,373536,373537,373538,373539,373540,373541,373542,373543,373544,373545,373546,373547,373548,373549,373550,373551,373552,373553,373554,373555,373556,373557,373558,373559,373560,373561,373562,373563,373564,373565,373566,373567,373568,373569,373570,373571,373572,373573,373574,373575,373576,373577,373578,373579,373580,373581,373582,373583,373584,373585,373586,373587,373588,373589,373590,373591,373592,373593,373594,373595,373596,373597,373598,373599,373600,373601,373602,373603,373604,373605,373606,373607,373608,373609,373610,373611,373612,373613,373614,373615,373616,373617,373618,373619,373620,373621,373622,373623,373624,373625,373626,373627,373628,373629,373630,373631,373632,373633,373634,373635,373636,373637,373638,373639,373640,373641,373642,373643,373644,373645,373646,373647,373648,373649,373650,373651,373652,373653,373654,373655,373656,373657,373658,373659,373660,373661,373662,373663,373664,373665,373666,373667,373668,373669,373670,373671,373672,373673,373674,373675,373676,373677,373678,373679,373680,373681,373682,373683,373684,373685,373686,373687,373688,373689,373690,373691,373692,373693,373694,373695,373696,373697,373698,373699,373700,373701,373702,373703,373704,373705,373706,373707,373708,373709,373710,373711,373712,373713,373714,373715,373716,373717,373718,373719,373720,373721,373722,373723,373724,373725,373726,373727,373728,373729,373730,373731,373732,373733,373734,373735,373736,373737,373738,373739,373740,373741,373742,373743,373744,373745,373746,373747,373748,373749,373750,373751,373752,373753,373754,373755,373756,373757,373758,373759,373760,373761,373762,373763,373764,373765,373766,373767,373768,373769,373770,373771,373772,373773,373774,373775,373776,373777,373778,373779,373780,373781,373782,373783,373784,373785,373786,373787,373788,373789,373790,373791,373792,373793,373794,373795,373796,373797,373798,373799,373800,373801,373802,373803,373804,373805,373806,373807,373808,373809,373810,373811,373812,373813,373814,373815,373816,373817,373818,373819,373820,373821,373822,373823,373824,373825,373826,373827,373828,373829,373830,373831,373832,373833,373834,373835,373836,373837,373838,373839,373840,373841,373842,373843,373844,373845,373846,373847,373848,373849,373850,373851,373852,373853,373854,373855,373856,373857,373858,373859,373860,373861,373862,373863,373864,373865,373866,373867,373868,373869,373870,373871,373872,373873,373874,373875,373876,373877,373878,373879,373880,373881,373882,373883,373884,373885,373886,373887,373888,373889,373890,373891,373892,373893,373894,373895,373896,373897,373898,373899,373900,373901,373902,373903,373904,373905,373906,373907,373908,373909,373910,373911,373912,373913,373914,373915,373916,373917,373918,373919,373920,373921,373922,373923,373924,373925,373926,373927,373928,373929,373930,373931,373932,373933,373934,373935,373936,373937,373938,373939,373940,373941,373942,373943,373944,373945,373946,373947,373948,373949,373950,373951,373952,373953,373954,373955,373956,373957,373958,373959,373960,373961,373962,373963,373964,373965,373966,373967,373968,373969,373970,373971,373972,373973,373974,373975,373976,373977,373978,373979,373980,373981,373982,373983,373984,373985,373986,373987,373988,373989,373990,373991,373992,373993,373994,373995,373996,373997,373998,373999,374000,374001,374002,374003,374004,374005,374006,374007,374008,374009,374010,374011,374012,374013,374014,374015,374016,374017,374018,374019,374020,374021,374022,374023,374024,374025,374026,374027,374028,374029,374030,374031,374032,374033,374034,374035,374036,374037,374038,374039,374040,374041,374042,374043,374044,374045,374046,374047,374048,374049,374050,374051,374052,374053,374054,374055,374056,374057,374058,374059,374060,374061,374062,374063,374064,374065,374066,374067,374068,374069,374070,374071,374072,374073,374074,374075,374076,374077,374078,374079,374080,374081,374082,374083,374084,374085,374086,374087,374088,374089,374090,374091,374092,374093,374094,374095,374096,374097,374098,374099,374100,374101,374102,374103,374104,374105,374106,374107,374108,374109,374110,374111,374112,374113,374114,374115,374116,374117,374118,374119,374120,374121,374122,374123,374124,374125,374126,374127,374128,374129,374130,374131,374132,374133,374134,374135,374136,374137,374138,374139,374140,374141,374142,374143,374144,374145,374146,374147,374148,374149,374150,374151,374152,374153,374154,374155,374156,374157,374158,374159,374160,374161,374162,374163,374164,374165,374166,374167,374168,374169,374170,374171,374172,374173,374174,374175,374176,374177,374178,374179,374180,374181,374182,374183,374184,374185,374186,374187,374188,374189,374190,374191,374192,374193,374194,374195,374196,374197,374198,374199,374200,374201,374202,374203,374204,374205,374206,374207,374208,374209,374210,374211,374212,374213,374214,374215,374216,374217,374218,374219,374220,374221,374222,374223,374224,374225,374226,374227,374228,374229,374230,374231,374232,374233,374234,374235,374236,374237,374238,374239,374240,374241,374242,374243,374244,374245,374246,374247,374248,374249,374250,374251,374252,374253,374254,374255,374256,374257,374258,374259,374260,374261,374262,374263,374264,374265,374266,374267,374268,374269,374270,374271,374272,374273,374274,374275,374276,374277,374278,374279,374280,374281,374282,374283,374284,374285,374286,374287,374288,374289,374290,374291,374292,374293,374294,374295,374296,374297,374298,374299,374300,374301,374302,374303,374304,374305,374306,374307,374308,374309,374325,374330,374331,374332,374333,374334,374335,374336,374337,374338,374339,374340,374341,374342,374343,374344,374345,374346,374347,374348,374349,374350,374351,374352,374353,374354,374355,374356,374357,374358,374359,374360,374361,374362,374363,374364,374365,374366,374367,374368,374369,374370,374371,374372,374373,374374,374375,374376,374377,374378,374379,374380,374381,374382,374383,374384,374385,374386,374387,374388,374389,374390,374391,374392,374393,374394,374395,374396,374397,374398,374399,374400,374401,374402,374403,374404,374405,374406,374407,374408,374409,374410,374411,374412,374413,374414,374415,374416,374417,374418,374419,374420,374421,374422,374423,374424,374425,374426,374427,374428,374429,374430,374431,374432,374433,374434,374435,374436,374437,374438,374439,374440,374441,374442,374443,374444,374445,374446,374447,374448,374449,374450,374451,374452,374453,374454,374455,374456,374457,374458,374459,374460,374461,374462,374463,374464,374465,374466,374467,374468,374469,374470,374471,374472,374473,374474,374475,374476,374477,374478,374479,374480,374481,374482,374483,374484,374485,374486,374487,374488,374489,374490,374491,374492,374493,374494,374495,374496,374497,374498,374499,374500,374501,374502,374503,374504,374505,374506,374507,374508,374509,374510,374511,374512,374513,374514,374515,374516,374517,374518,374519,374520,374521,374522,374523,374524,374525,374526,374527,374528,374529,374530,374531,374532,374533,374534,374535,374536,374537,374538,374539,374540,374541,374542,374543,374544,374545,374546,374547,374548,374549,374550,374551,374552,374553,374554,374555,374556,374557,374558,374559,374560,374561,374562,374563,374564,374565,374566,374567,374568,374569,374570,374571,374572,374573,374574,374575,374576,374577,374578,374579,374580,374581,374582,374583,374584,374585,374586,374587,374588,374589,374590,374591,374592,374593,374594,374595,374596,374597,374598,374599,374600,374601,374602,374603,374604,374605,374606,374607,374608,374609,374610,374611,374612,374613,374614,374615,374616,374617,374618,374619,374620,374621,374622,374623,374624,374625,374626,374627,374628,374629,374640,374641,374642,374643,374644,374645,374646,374647,374648,374649,374650,374651,374652,374653,374654,374655,374656,374657,374658,374659,374660,374661,374662,374663,374664,374665,374666,374667,374668,374669,374670,374671,374672,374673,374674,374675,374676,374677,374678,374679,374680,374681,374682,374683,374684,374685,374686,374687,374688,374689,374690,374691,374692,374693,374694,374695,374696,374697,374698,374699,374700,374701,374702,374703,374704,374705,374706,374707,374708,374709,374710,374711,374722,374730,374731,374732,374733,374734,374735,374736,374737,374738,374739,374740,374741,374742,374743,374744,374745,374746,374747,374748,374749,374750,374751,374752,374753,374754,374755,374756,374757,374758,374759,374760,374761,374762,374763,374764,374765,374766,374767,374768,374769,374770,374771,374772,374773,374774,374775,374776,374777,374778,374779,374780,374781,374782,374783,374784,374785,374786,374787,374788,374789,374790,374791,374792,374793,374794,374795,374796,374797,374798,374799,374800,374801,374802,374803,374804,374805,374806,374807,374808,374809,374810,374811,374812,374813,374814,374815,374816,374817,374818,374819,374820,374821,374822,374823,374824,374825,374826,374827,374828,374829,374830,374831,374832,374833,374834,374835,374836,374837,374838,374839,374840,374841,374842,374843,374844,374845,374846,374847,374848,374849,374850,374851,374852,374853,374854,374855,374856,374857,374858,374859,374860,374861,374862,374863,374864,374865,374866,374867,374868,374869,374870,374871,374872,374873,374874,374875,374876,374877,374878,374879,374880,374881,374882,374883,374884,374885,374886,374887,374888,374889,374890,374891,374892,374893,374894,374895,374896,374897,374898,374899,374900,374901,374902,374903,374904,374905,374906,374907,374908,374909,374910,374911,374912,374913,374914,374915,374916,374917,374918,374919,374920,374921,374922,374923,374924,374925,374926,374927,374928,374929,374930,374931,374932,374933,374934,374935,374936,374937,374938,374939,374940,374941,374942,374943,374944,374945,374946,374947,374948,374949,374950,374951,374952,374953,374954,374955,374956,374957,374958,374959,374960,374961,374962,374963,374964,374965,374966,374967,374968,374969,374970,374971,374972,374973,374974,374975,374976,374977,374978,374979,374980,374981,374982,374983,374984,374985,374986,374987,374988,374989,374990,374991,374992,374993,374994,374995,374996,374997,374998,374999,375000,375001,375002,375003,375004,375005,375006,375007,375008,375009,375010,375011,375012,375013,375014,375015,375016,375017,375018,375019,375020,375021,375022,375023,375024,375025,375026,375027,375028,375029,375030,375031,375032,375033,375034,375035,375036,375037,375038,375039,375040,375041,375042,375043,375044,375045,375046,375047,375048,375049,375050,375051,375052,375053,375054,375055,375056,375057,375058,375059,375060,375061,375062,375063,375064,375065,375066,375067,375068,375069,375070,375071,375072,375073,375074,375075,375076,375077,375078,375079,375080,375081,375082,375083,375084,375085,375086,375087,375088,375089,375090,375091,375092,375093,375094,375095,375096,375097,375098,375099,375100,375101,375102,375103,375104,375105,375106,375107,375108,375109,375110,375111,375112,375113,375114,375115,375116,375117,375118,375119,375120,375121,375122,375123,375124,375125,375126,375127,375128,375129,375130,375131,375132,375133,375134,375135,375136,375137,375138,375139,375140,375141,375142,375143,375144,375145,375146,375147,375148,375149,375150,375151,375152,375153,375154,375155,375156,375157,375158,375159,375160,375161,375162,375163,375164,375165,375166,375167,375168,375169,375170,375171,375172,375173,375174,375175,375176,375177,375178,375179,375180,375181,375182,375183,375184,375185,375186,375187,375188,375189,375190,375191,375192,375193,375194,375195,375196,375197,375198,375199,375200,375201,375202,375203,375204,375205,375206,375207,375208,375209,375210,375211,375212,375213,375214,375215,375216,375217,375218,375219,375220,375221,375222,375223,375224,375225,375226,375227,375228,375229,375230,375231,375232,375233,375234,375235,375236,375237,375238,375239,375240,375241,375242,375243,375244,375245,375246,375247,375248,375249,375250,375251,375252,375253,375254,375255,375256,375257,375258,375259,375260,375261,375262,375263,375264,375265,375266,375267,375268,375269,375270,375271,375272,375273,375274,375275,375276,375277,375278,375279,375280,375281,375282,375283,375284,375285,375286,375287,375288,375289,375290,375291,375292,375293,375294,375295,375296,375297,375298,375299,375300,375301,375302,375303,375304,375305,375306,375307,375308,375309,375310,375311,375312,375313,375314,375315,375316,375317,375318,375319,375320,375321,375322,375323,375324,375325,375326,375327,375328,375329,375330,375331,375332,375333,375334,375335,375336,375337,375338,375339,375340,375341,375342,375343,375344,375345,375346,375347,375348,375349,375350,375351,375352,375353,375354,375355,375356,375357,375358,375359,375360,375361,375362,375363,375364,375365,375366,375367,375368,375369,375370,375371,375372,375373,375374,375375,375376,375377,375378,375379,375380,375381,375382,375383,375384,375385,375386,375387,375388,375389,375390,375391,375392,375393,375394,375395,375396,375397,375398,375399,375400,375401,375402,375403,375404,375405,375406,375407,375408,375409,375410,375411,375412,375413,375414,375415,375416,375417,375418,375419,375420,375421,375422,375423,375424,375425,375426,375427,375428,375429,375430,375431,375432,375433,375434,375435,375436,375437,375438,375439,375440,375441,375442,375443,375444,375445,375446,375447,375448,375449,375450,375451,375452,375453,375454,375455,375456,375457,375458,375459,375460,375461,375462,375463,375464,375465,375466,375467,375468,375469,375470,375471,375472,375473,375474,375475,375476,375477,375478,375479,375480,375481,375482,375483,375484,375485,375486,375487,375488,375489,375490,375491,375492,375493,375494,375495,375496,375497,375498,375499,375500,375501,375502,375503,375504,375505,375506,375507,375508,375509,375510,375511,375512,375513,375514,375515,375516,375517,375518,375519,375520,375521,375522,375523,375524,375525,375526,375527,375528,375529,375530,375531,375532,375533,375534,375535,375536,375537,375538,375539,375540,375541,375542,375543,375544,375545,375546,375547,375548,375549,375550,375551,375552,375553,375554,375555,375556,375557,375558,375559,375560,375561,375562,375563,375564,375565,375566,375567,375568,375569,375570,375571,375572,375573,375574,375575,375576,375577,375578,375579,375580,375581,375582,375583,375584,375585,375586,375587,375588,375589,375590,375591,375592,375593,375594,375595,375596,375597,375598,375599,375600,375601,375602,375603,375604,375605,375606,375607,375608,375609,375610,375611,375612,375613,375614,375615,375616,375617,375618,375619,375620,375621,375622,375623,375624,375625,375626,375627,375628,375629,375630,375631,375632,375633,375634,375635,375636,375637,375638,375639,375640,375641,375642,375643,375644,375645,375646,375647,375648,375649,375650,375651,375652,375653,375654,375655,375656,375657,375658,375659,375660,375661,375662,375663,375664,375665,375666,375667,375668,375669,375670,375671,375672,375673,375674,375675,375676,375677,375678,375679,375680,375681,375682,375683,375684,375685,375686,375687,375688,375689,375690,375691,375692,375693,375694,375695,375696,375697,375698,375699,375700,375701,375702,375703,375704,375705,375706,375707,375708,375709,375710,375711,375712,375713,375714,375715,375716,375717,375718,375719,375720,375721,375722,375723,375724,375725,375726,375727,375728,375729,375730,375731,375732,375733,375734,375735,375736,375737,375738,375739,375740,375741,375742,375743,375744,375745,375746,375747,375748,375749,375750,375751,375752,375753,375754,375755,375756,375757,375758,375759,375760,375761,375762,375763,375764,375765,375766,375767,375768,375769,375770,375771,375772,375773,375774,375775,375776,375777,375778,375779,375780,375781,375782,375783,375784,375785,375786,375787,375788,375789,375790,375791,375792,375793,375794,375795,375796,375797,375798,375799,375800,375801,375802,375803,375804,375805,375806,375807,375808,375809,375810,375811,375812,375813,375814,375815,375816,375817,375818,375819,375820,375821,375822,375823,375824,375825,375826,375827,375828,375829,375830,375831,375832,375833,375834,375835,375836,375837,375838,375839,375840,375841,375842,375843,375844,375845,375846,375847,375848,375849,375850,375851,375852,375853,375854,375855,375856,375857,375858,375859,375860,375861,375862,375863,375864,375865,375866,375867,375868,375869,375870,375871,375872,375873,375874,375875,375876,375877,375878,375879,375880,375881,375882,375883,375884,375885,375886,375887,375888,375889,375890,375891,375892,375893,375894,375895,375896,375897,375898,375899,375900,375901,375902,375903,375904,375905,375906,375907,375908,375909,375910,375911,375912,375913,375914,375915,375916,375917,375918,375919,375920,375921,375922,375923,375924,375925,375926,375927,375928,375929,375930,375931,375932,375933,375934,375935,375936,375937,375938,375939,375940,375941,375942,375943,375944,375945,375946,375947,375948,375949,375950,375951,375952,375953,375954,375955,375956,375957,375958,375959,375960,375961,375962,375963,375964,375965,375966,375967,375968,375969,375970,375971,375972,375973,375974,375975,375976,375977,375978,375979,375980,375981,375982,375983,375984,375985,375986,375987,375988,375989,375990,375991,375992,375993,375994,375995,375996,375997,375998,375999,376000,376001,376002,376003,376004,376005,376006,376007,376008,376009,376010,376011,376012,376013,376014,376015,376016,376017,376018,376019,376020,376021,376022,376023,376024,376025,376026,376027,376028,376029,376030,376031,376032,376033,376034,376035,376036,376037,376038,376039,376040,376041,376042,376043,376044,376045,376046,376047,376048,376049,376050,376051,376052,376053,376054,376055,376056,376057,376058,376059,376060,376061,376062,376063,376064,376065,376066,376067,376068,376069,376070,376071,376072,376073,376074,376075,376076,376077,376078,376079,376080,376081,376082,376083,376084,376085,376086,376087,376088,376089,376090,376091,376092,376093,376094,376095,376096,376097,376098,376099,376100,376101,376102,376103,376104,376105,376106,376107,376108,376109,376110,376111,376112,376113,376114,376115,376116,376117,376118,376119,376120,376121,376122,376123,376124,376125,376126,376127,376128,376129,376130,376131,376132,376133,376134,376135,376136,376137,376138,376139,376140,376141,376142,376143,376144,376145,376146,376147,376148,376149,376150,376151,376152,376153,376154,376155,376156,376157,376158,376159,376160,376161,376162,376163,376164,376165,376166,376167,376168,376169,376170,376171,376172,376173,376174,376175,376176,376177,376178,376179,376180,376181,376182,376183,376184,376185,376186,376187,376188,376189,376190,376191,376192,376193,376194,376195,376196,376197,376198,376199,376200,376201,376202,376203,376204,376205,376206,376207,376208,376209,376210,376211,376212,376213,376214,376215,376216,376217,376218,376219,376220,376221,376222,376223,376224,376225,376226,376227,376228,376229,376230,376231,376232,376233,376234,376235,376236,376237,376238,376239,376240,376241,376242,376243,376244,376245,376246,376247,376248,376249,376250,376251,376252,376253,376254,376255,376256,376257,376258,376259,376260,376261,376262,376263,376264,376265,376266,376267,376268,376269,376270,376271,376272,376273,376274,376275,376276,376277,376278,376279,376280,376281,376282,376283,376284,376285,376286,376287,376288,376289,376290,376291,376292,376293,376294,376295,376296,376297,376298,376299,376300,376301,376302,376303,376304,376305,376306,376307,376308,376309,376310,376311,376312,376313,376314,376315,376316,376317,376318,376319,376320,376321,376322,376323,376324,376325,376326,376327,376328,376329,376330,376331,376332,376333,376334,376335,376336,376337,376338,376339,376340,376341,376342,376343,376344,376345,376346,376347,376348,376349,376350,376351,376352,376353,376354,376355,376356,376357,376358,376359,376360,376361,376362,376363,376364,376365,376366,376367,376368,376369,376370,376371,376372,376373,376374,376375,376376,376377,376378,376379,376380,376381,376382,376383,376384,376385,376386,376387,376388,376389,376390,376391,376392,376393,376394,376395,376396,376397,376398,376399,376400,376401,376402,376403,376404,376405,376406,376407,376408,376409,376410,376411,376412,376413,376414,376415,376416,376417,376418,376419,376420,376421,376422,376423,376424,376425,376426,376427,376428,376429,376430,376431,376432,376433,376434,376435,376436,376437,376438,376439,376440,376441,376442,376443,376444,376445,376446,376447,376448,376449,376450,376451,376452,376453,376454,376455,376456,376457,376458,376459,376460,376461,376462,376463,376464,376465,376466,376467,376468,376469,376470,376471,376472,376473,376474,376475,376476,376477,376478,376479,376480,376481,376482,376483,376484,376485,376486,376487,376488,376489,376490,376491,376492,376493,376494,376495,376496,376497,376498,376499,376500,376501,376502,376503,376504,376505,376506,376507,376508,376509,376510,376511,376512,376513,376514,376515,376516,376517,376518,376519,376520,376521,376522,376523,376524,376525,376526,376527,376528,376529,376530,376531,376532,376533,376534,376535,376536,376537,376538,376539,376540,376541,376542,376543,376544,376545,376546,376547,376548,376549,376550,376551,376552,376553,376554,376555,376556,376557,376558,376559,376560,376561,376562,376563,376564,376565,376566,376567,376568,376569,376570,376571,376572,376573,376574,376575,376576,376577,376578,376579,376580,376581,376582,376583,376584,376585,376586,376587,376588,376589,376590,376591,376592,376593,376594,376595,376596,376597,376598,376599,376600,376601,376602,376603,376604,376605,376606,376607,376608,376609,376610,376611,376612,376613,376614,376615,376616,376617,376618,376619,376620,376621,376622,376623,376624,376625,376626,376627,376628,376629,376630,376631,376632,376633,376634,376635,376636,376637,376638,376639,376640,376641,376642,376643,376644,376645,376646,376647,376648,376649,376650,376651,376652,376653,376654,376655,376656,376657,376658,376659,376660,376661,376662,376663,376664,376665,376666,376667,376668,376669,376670,376671,376672,376673,376674,376675,376676,376677,376678,376679,376680,376681,376682,376683,376684,376685,376686,376687,376688,376689,376690,376691,376692,376693,376694,376695,376696,376697,376698,376699,376700,376701,376702,376703,376704,376705,376706,376707,376708,376709,376710,376711,376712,376713,376714,376715,376716,376717,376718,376719,376720,376721,376722,376723,376724,376725,376726,376727,376728,376729,376730,376731,376732,376733,376734,376735,376736,376737,376738,376739,376740,376741,376742,376743,376744,376745,376746,376747,376748,376749,376750,376751,376752,376753,376754,376755,376756,376757,376758,376759,376760,376761,376762,376763,376764,376765,376766,376767,376768,376769,376770,376771,376772,376773,376774,376775,376776,376777,376778,376779,376780,376781,376782,376783,376784,376785,376786,376787,376788,376789,376790,376791,376792,376793,376794,376795,376796,376797,376798,376799,376800,376801,376802,376803,376804,376805,376806,376807,376808,376809,376810,376811,376812,376813,376814,376815,376816,376817,376818,376819,376820,376821,376822,376823,376824,376825,376826,376827,376828,376829,376830,376831,376832,376833,376834,376835,376836,376837,376838,376839,376840,376841,376842,376843,376844,376845,376846,376847,376848,376849,376850,376851,376852,376853,376854,376855,376856,376857,376858,376859,376860,376861,376862,376863,376864,376865,376866,376867,376868,376869,376870,376871,376872,376873,376874,376875,376876,376877,376878,376879,376880,376881,376882,376883,376884,376885,376886,376887,376888,376889,376890,376891,376892,376893,376894,376895,376896,376897,376898,376899,376900,376901,376902,376903,376904,376905,376906,376907,376908,376909,376910,376911,376912,376913,376914,376915,376916,376917,376918,376919,376920,376921,376922,376923,376924,376925,376926,376927,376928,376929,376930,376931,376932,376933,376934,376935,376936,376937,376938,376939,376940,376941,376942,376943,376944,376945,376946,376947,376948,376949,376950,376951,376952,376953,376954,376955,376956,376957,376958,376959,376960,376961,376962,376963,376964,376965,376966,376967,376968,376969,376970,376971,376972,376973,376974,376975,376976,376977,376978,376979,376980,376981,376982,376983,376984,376985,376986,376987,376988,376989,376990,376991,376992,376993,376994,376995,376996,376997,376998,376999,377000,377001,377002,377003,377004,377005,377006,377007,377008,377009,377010,377011,377012,377013,377014,377015,377016,377017,377018,377019,377020,377021,377022,377023,377024,377025,377026,377027,377028,377029,377052,377053,377054,377055,377056,377057,377058,377059,377100,377101,377102,377103,377104,377105,377106,377107,377108,377109,377110,377111,377112,377113,377114,377115,377116,377117,377118,377119,377120,377121,377122,377123,377124,377125,377126,377127,377128,377129,377130,377131,377132,377133,377134,377135,377136,377137,377138,377139,377140,377141,377142,377143,377144,377145,377146,377147,377148,377149,377150,377151,377152,377153,377154,377155,377156,377157,377158,377159,377160,377161,377162,377163,377164,377165,377166,377167,377168,377169,377170,377171,377172,377173,377174,377175,377176,377177,377178,377179,377180,377181,377182,377183,377184,377185,377186,377187,377188,377189,377190,377191,377192,377193,377194,377195,377196,377197,377198,377199,377200,377201,377202,377203,377204,377205,377206,377207,377208,377209,377210,377211,377212,377213,377214,377215,377216,377217,377218,377219,377220,377221,377222,377223,377224,377225,377226,377227,377228,377229,377230,377231,377232,377233,377234,377235,377236,377237,377238,377239,377240,377241,377242,377243,377244,377245,377246,377247,377248,377249,377250,377251,377252,377253,377254,377255,377256,377257,377258,377259,377260,377261,377262,377263,377264,377265,377266,377267,377268,377269,377270,377271,377272,377273,377274,377275,377276,377277,377278,377279,377280,377281,377282,377283,377284,377285,377286,377287,377288,377289,377290,377291,377292,377293,377294,377295,377296,377297,377298,377299,377300,377301,377302,377303,377304,377305,377306,377307,377308,377309,377310,377311,377312,377313,377314,377315,377316,377317,377318,377319,377320,377321,377322,377323,377324,377325,377326,377327,377328,377329,377330,377331,377332,377333,377334,377335,377336,377337,377338,377339,377340,377341,377342,377343,377344,377345,377346,377347,377348,377349,377350,377351,377352,377353,377354,377355,377356,377357,377358,377359,377360,377361,377362,377363,377364,377365,377366,377367,377368,377369,377370,377371,377372,377373,377374,377375,377376,377377,377378,377379,377380,377381,377382,377383,377384,377385,377386,377387,377388,377389,377390,377391,377392,377393,377394,377395,377396,377397,377398,377399,377400,377401,377402,377403,377404,377405,377406,377407,377408,377409,377410,377411,377412,377413,377414,377415,377416,377417,377418,377419,377420,377421,377422,377423,377424,377425,377426,377427,377428,377429,377430,377431,377432,377433,377434,377435,377436,377437,377438,377439,377440,377441,377442,377443,377444,377445,377446,377447,377448,377449,377450,377451,377452,377453,377454,377455,377456,377457,377458,377459,377470,377471,377472,377473,377474,377475,377476,377477,377478,377479,377480,377481,377482,377483,377484,377485,377486,377487,377488,377489,377490,377491,377492,377493,377494,377495,377496,377497,377498,377499,377500,377501,377502,377503,377504,377505,377506,377507,377508,377509,377510,377511,377512,377513,377514,377515,377516,377517,377518,377519,377520,377521,377522,377523,377524,377525,377526,377527,377528,377529,377530,377531,377532,377533,377534,377535,377536,377537,377538,377539,377540,377541,377542,377543,377544,377545,377546,377547,377548,377549,377550,377551,377552,377553,377554,377555,377556,377557,377558,377559,377560,377561,377562,377563,377564,377565,377566,377567,377568,377569,377570,377571,377572,377573,377574,377575,377576,377577,377578,377579,377580,377581,377582,377583,377584,377585,377586,377587,377588,377589,377590,377591,377592,377593,377594,377595,377596,377597,377598,377599,377600,377601,377602,377603,377604,377605,377606,377607,377608,377609,377610,377611,377612,377613,377614,377615,377616,377617,377618,377619,377620,377621,377622,377623,377624,377625,377626,377627,377628,377629,377630,377631,377632,377633,377634,377635,377636,377637,377638,377639,377640,377641,377642,377643,377644,377645,377646,377647,377648,377649,377650,377651,377652,377653,377654,377655,377656,377657,377658,377659,377660,377661,377662,377663,377664,377665,377666,377667,377668,377669,377670,377671,377672,377673,377674,377675,377676,377677,377678,377679,377680,377681,377682,377683,377684,377685,377686,377687,377688,377689,377690,377691,377692,377693,377694,377695,377696,377697,377698,377699,377700,377701,377702,377703,377704,377705,377706,377707,377708,377709,377710,377711,377712,377713,377714,377715,377716,377717,377718,377719,377720,377721,377722,377723,377724,377725,377726,377727,377728,377729,377730,377731,377732,377733,377734,377735,377736,377737,377738,377739,377740,377741,377742,377743,377744,377745,377746,377747,377748,377749,377750,377751,377752,377753,377754,377755,377756,377757,377758,377759,377760,377761,377762,377763,377764,377765,377766,377767,377768,377769,377770,377771,377772,377773,377774,377775,377776,377777,377778,377779,377780,377781,377782,377783,377784,377785,377786,377787,377788,377789,377790,377791,377792,377793,377794,377795,377796,377797,377798,377799,377800,377801,377802,377803,377804,377805,377806,377807,377808,377809,377820,377821,377822,377823,377824,377825,377826,377827,377828,377829,377830,377831,377832,377833,377834,377835,377836,377837,377838,377839,377840,377841,377842,377843,377844,377845,377846,377847,377848,377849,377850,377851,377852,377853,377854,377855,377856,377857,377858,377859,377860,377861,377862,377863,377864,377865,377866,377867,377868,377869,377870,377871,377872,377873,377874,377875,377876,377877,377878,377879,377880,377881,377882,377883,377884,377885,377886,377887,377888,377889,377890,377891,377892,377893,377894,377895,377896,377897,377898,377899,378000,378001,378002,378003,378004,378005,378006,378007,378008,378009,378010,378011,378012,378013,378014,378015,378016,378017,378018,378019,378020,378021,378022,378023,378024,378025,378026,378027,378028,378029,378030,378031,378032,378033,378034,378035,378036,378037,378038,378039,378040,378041,378042,378043,378044,378045,378046,378047,378048,378049,378050,378051,378052,378053,378054,378055,378056,378057,378058,378059,378060,378061,378062,378063,378064,378065,378066,378067,378068,378069,378070,378071,378072,378073,378074,378075,378076,378077,378078,378079,378080,378081,378082,378083,378084,378085,378086,378087,378088,378089,378090,378091,378092,378093,378094,378095,378096,378097,378098,378099,378100,378101,378102,378103,378104,378105,378106,378107,378108,378109,378110,378111,378112,378113,378114,378115,378116,378117,378118,378119,378120,378121,378122,378123,378124,378125,378126,378127,378128,378129,378130,378131,378132,378133,378134,378135,378136,378137,378138,378139,378140,378141,378142,378143,378144,378145,378146,378147,378148,378149,378150,378151,378152,378153,378154,378155,378156,378157,378158,378159,378160,378161,378162,378163,378164,378165,378166,378167,378168,378169,378170,378171,378172,378173,378174,378175,378176,378177,378178,378179,378180,378181,378182,378183,378184,378185,378186,378187,378188,378189,378190,378191,378192,378193,378194,378195,378196,378197,378198,378199,378200,378201,378202,378203,378204,378205,378206,378207,378208,378209,378210,378211,378212,378213,378214,378215,378216,378217,378218,378219,378220,378221,378222,378223,378224,378225,378226,378227,378228,378229,378230,378231,378232,378233,378234,378235,378236,378237,378238,378239,378240,378241,378242,378243,378244,378245,378246,378247,378248,378249,378250,378251,378252,378253,378254,378255,378256,378257,378258,378259,378260,378261,378262,378263,378264,378265,378266,378267,378268,378269,378270,378271,378272,378273,378274,378275,378276,378277,378278,378279,378280,378281,378282,378283,378284,378285,378286,378287,378288,378289,378290,378291,378292,378293,378294,378295,378296,378297,378298,378299,378300,378301,378302,378303,378304,378305,378306,378307,378308,378309,378310,378311,378312,378313,378314,378315,378316,378317,378318,378319,378320,378321,378322,378323,378324,378325,378326,378327,378328,378329,378330,378331,378332,378333,378334,378335,378336,378337,378338,378339,378340,378341,378342,378343,378344,378345,378346,378347,378348,378349,378350,378351,378352,378353,378354,378355,378356,378357,378358,378359,378360,378361,378362,378363,378364,378365,378366,378367,378368,378369,378370,378371,378372,378373,378374,378375,378376,378377,378378,378379,378380,378381,378382,378383,378384,378385,378386,378387,378388,378389,378390,378391,378392,378393,378394,378395,378396,378397,378398,378399,378400,378401,378402,378403,378404,378405,378406,378407,378408,378409,378410,378411,378412,378413,378414,378415,378416,378417,378418,378419,378420,378421,378422,378423,378424,378425,378426,378427,378428,378429,378430,378431,378432,378433,378434,378435,378436,378437,378438,378439,378440,378441,378442,378443,378444,378445,378446,378447,378448,378449,378450,378451,378452,378453,378454,378455,378456,378457,378458,378459,378460,378461,378462,378463,378464,378465,378466,378467,378468,378469,378470,378471,378472,378473,378474,378475,378476,378477,378478,378479,378480,378481,378482,378483,378484,378485,378486,378487,378488,378489,378490,378491,378492,378493,378494,378495,378496,378497,378498,378499,378500,378501,378502,378503,378504,378505,378506,378507,378508,378509,378510,378511,378512,378513,378514,378515,378516,378517,378518,378519,378520,378521,378522,378523,378524,378525,378526,378527,378528,378529,378530,378531,378532,378533,378534,378535,378536,378537,378538,378539,378540,378541,378542,378543,378544,378545,378546,378547,378548,378549,378550,378551,378552,378553,378554,378555,378556,378557,378558,378559,378560,378561,378562,378563,378564,378565,378566,378567,378568,378569,378570,378571,378572,378573,378574,378575,378576,378577,378578,378579,378580,378581,378582,378583,378584,378585,378586,378587,378588,378589,378590,378591,378592,378593,378594,378595,378596,378597,378598,378599,378600,378601,378602,378603,378604,378605,378606,378607,378608,378609,378610,378611,378612,378613,378614,378615,378616,378617,378618,378619,378620,378621,378622,378623,378624,378625,378626,378627,378628,378629,378630,378631,378632,378633,378634,378635,378636,378637,378638,378639,378640,378641,378642,378643,378644,378645,378646,378647,378648,378649,378650,378651,378652,378653,378654,378655,378656,378657,378658,378659,378660,378661,378662,378663,378664,378665,378666,378667,378668,378669,378670,378671,378672,378673,378674,378675,378676,378677,378678,378679,378680,378681,378682,378683,378684,378685,378686,378687,378688,378689,378690,378691,378692,378693,378694,378695,378696,378697,378698,378699,378700,378701,378702,378703,378704,378705,378706,378707,378708,378709,378710,378711,378712,378713,378714,378715,378716,378717,378718,378719,378720,378721,378722,378723,378724,378725,378726,378727,378728,378729,378730,378731,378732,378733,378734,378735,378736,378737,378738,378739,378740,378741,378742,378743,378744,378745,378746,378747,378748,378749,378750,378751,378752,378753,378754,378755,378756,378757,378758,378759,378760,378761,378762,378763,378764,378765,378766,378767,378768,378769,378770,378771,378772,378773,378774,378775,378776,378777,378778,378779,378780,378781,378782,378783,378784,378785,378786,378787,378788,378789,378790,378791,378792,378793,378794,378795,378796,378797,378798,378799,378800,378801,378802,378803,378804,378805,378806,378807,378808,378809,378810,378811,378812,378813,378814,378815,378816,378817,378818,378819,378820,378821,378822,378823,378824,378825,378826,378827,378828,378829,378830,378831,378832,378833,378834,378835,378836,378837,378838,378839,378840,378841,378842,378843,378844,378845,378846,378847,378848,378849,378850,378851,378852,378853,378854,378855,378856,378857,378858,378859,378860,378861,378862,378863,378864,378865,378866,378867,378868,378869,378870,378871,378872,378873,378874,378875,378876,378877,378878,378879,378880,378881,378882,378883,378884,378885,378886,378887,378888,378889,378890,378891,378892,378893,378894,378895,378896,378897,378898,378899,378900,378901,378902,378903,378904,378905,378906,378907,378908,378909,378910,378911,378912,378913,378914,378915,378916,378917,378918,378919,378920,378921,378922,378923,378924,378925,378926,378927,378928,378929,378930,378931,378932,378933,378934,378935,378936,378937,378938,378939,378940,378941,378942,378943,378944,378945,378946,378947,378948,378949,378950,378951,378952,378953,378954,378955,378956,378957,378958,378959,378960,378961,378962,378963,378964,378965,378966,378967,378968,378969,378970,378971,378972,378973,378974,378975,378976,378977,378978,378979,378980,378981,378982,378983,378984,378985,378986,378987,378988,378989,378990,378991,378992,378993,378994,378995,378996,378997,378998,378999,379000,379001,379002,379003,379004,379005,379006,379007,379008,379009,379010,379011,379012,379013,379014,379015,379016,379017,379018,379019,379020,379021,379022,379023,379024,379025,379026,379027,379028,379029,379030,379031,379032,379033,379034,379035,379036,379037,379038,379039,379040,379041,379042,379043,379044,379045,379046,379047,379048,379049,379050,379051,379052,379053,379054,379055,379056,379057,379058,379059,379060,379061,379062,379063,379064,379065,379066,379067,379068,379069,379070,379071,379072,379073,379074,379075,379076,379077,379078,379079,379080,379081,379082,379083,379084,379085,379086,379087,379088,379089,379090,379091,379092,379093,379094,379095,379096,379097,379098,379099,379100,379101,379102,379103,379104,379105,379106,379107,379108,379109,379110,379111,379112,379113,379114,379115,379116,379117,379118,379119,379120,379121,379122,379123,379124,379125,379126,379127,379128,379129,379130,379131,379132,379133,379134,379135,379136,379137,379138,379139,379140,379141,379142,379143,379144,379145,379146,379147,379148,379149,379150,379151,379152,379153,379154,379155,379156,379157,379158,379159,379160,379161,379162,379163,379164,379165,379166,379167,379168,379169,379170,379171,379172,379173,379174,379175,379176,379177,379178,379179,379180,379181,379182,379183,379184,379185,379186,379187,379188,379189,379190,379191,379192,379193,379194,379195,379196,379197,379198,379199,379200,379201,379202,379203,379204,379205,379206,379207,379208,379209,379210,379211,379212,379213,379214,379215,379216,379217,379218,379219,379220,379221,379222,379223,379224,379225,379226,379227,379228,379229,379230,379231,379232,379233,379234,379235,379236,379237,379238,379239,379240,379241,379242,379243,379244,379245,379246,379247,379248,379249,379250,379251,379252,379253,379254,379255,379256,379257,379258,379259,379260,379261,379262,379263,379264,379265,379266,379267,379268,379269,379270,379271,379272,379273,379274,379275,379276,379277,379278,379279,379280,379281,379282,379283,379284,379285,379286,379287,379288,379289,379290,379291,379292,379293,379294,379295,379296,379297,379298,379299,379300,379301,379302,379303,379304,379305,379306,379307,379308,379309,379310,379311,379312,379313,379314,379315,379316,379317,379318,379319,379320,379321,379322,379323,379324,379325,379326,379327,379328,379329,379330,379331,379332,379333,379334,379335,379336,379337,379338,379339,379340,379341,379342,379343,379344,379345,379346,379347,379348,379349,379350,379351,379352,379353,379354,379355,379356,379357,379358,379359,379360,379361,379362,379363,379364,379365,379366,379367,379368,379369,379370,379371,379372,379373,379374,379375,379376,379377,379378,379379,379380,379381,379382,379383,379384,379385,379386,379387,379388,379389,379390,379391,379392,379393,379394,379395,379396,379397,379398,379399,379400,379401,379402,379403,379404,379405,379406,379407,379408,379409,379410,379411,379412,379413,379414,379415,379416,379417,379418,379419,379420,379421,379422,379423,379424,379425,379426,379427,379428,379429,379430,379431,379432,379433,379434,379435,379436,379437,379438,379439,379440,379441,379442,379443,379444,379445,379446,379447,379448,379449,379450,379451,379452,379453,379454,379455,379456,379457,379458,379459,379460,379461,379462,379463,379464,379465,379466,379467,379468,379469,379470,379471,379472,379473,379474,379475,379476,379477,379478,379479,379480,379481,379482,379483,379484,379485,379486,379487,379488,379489,379490,379491,379492,379493,379494,379495,379496,379497,379498,379499,379500,379501,379502,379503,379504,379505,379506,379507,379508,379509,379510,379511,379512,379513,379514,379515,379516,379517,379518,379519,379520,379521,379522,379523,379524,379525,379526,379527,379528,379529,379530,379531,379532,379533,379534,379535,379536,379537,379538,379539,379540,379541,379542,379543,379544,379545,379546,379547,379548,379549,379550,379551,379552,379553,379554,379555,379556,379557,379558,379559,379560,379561,379562,379563,379564,379565,379566,379567,379568,379569,379570,379571,379572,379573,379574,379575,379576,379577,379578,379579,379580,379581,379582,379583,379584,379585,379586,379587,379588,379589,379590,379591,379592,379593,379594,379595,379596,379597,379598,379599,379600,379601,379602,379603,379604,379605,379606,379607,379608,379609,379610,379611,379612,379613,379614,379615,379616,379617,379618,379619,379620,379621,379622,379623,379624,379625,379626,379627,379628,379629,379630,379631,379632,379633,379634,379635,379636,379637,379638,379639,379640,379641,379642,379643,379644,379645,379646,379647,379648,379649,379650,379651,379652,379653,379654,379655,379656,379657,379658,379659,379660,379661,379662,379663,379664,379665,379666,379667,379668,379669,379670,379671,379672,379673,379674,379675,379676,379677,379678,379679,379680,379681,379682,379683,379684,379685,379686,379687,379688,379689,379690,379691,379692,379693,379694,379695,379696,379697,379698,379699,379700,379701,379702,379703,379704,379705,379706,379707,379708,379709,379710,379711,379712,379713,379714,379715,379716,379717,379718,379719,379720,379721,379722,379723,379724,379725,379726,379727,379728,379729,379730,379731,379732,379733,379734,379735,379736,379737,379738,379739,379740,379741,379742,379743,379744,379745,379746,379747,379748,379749,379750,379751,379752,379753,379754,379755,379756,379757,379758,379759,379760,379761,379762,379763,379764,379765,379766,379767,379768,379769,379770,379771,379772,379773,379774,379775,379776,379777,379778,379779,379780,379781,379782,379783,379784,379785,379786,379787,379788,379789,379790,379791,379792,379793,379794,379795,379796,379797,379798,379799,379800,379801,379802,379803,379804,379805,379806,379807,379808,379809,379810,379811,379812,379813,379814,379815,379816,379817,379818,379819,379820,379821,379822,379823,379824,379825,379826,379827,379828,379829,379830,379831,379832,379833,379834,379835,379836,379837,379838,379839,379840,379841,379842,379843,379844,379845,379846,379847,379848,379849,379850,379851,379852,379853,379854,379855,379856,379857,379858,379859,379860,379861,379862,379863,379864,379865,379866,379867,379868,379869,379870,379871,379872,379873,379874,379875,379876,379877,379878,379879,379880,379881,379882,379883,379884,379885,379886,379887,379888,379889,379890,379891,379892,379893,379894,379895,379896,379897,379898,379899,379900,379901,379902,379903,379904,379905,379906,379907,379908,379909,379910,379911,379912,379913,379914,379915,379916,379917,379918,379919,379920,379921,379922,379923,379924,379925,379926,379927,379928,379929,379930,379931,379932,379933,379934,379935,379936,379937,379938,379939,379940,379941,379942,379943,379944,379945,379946,379947,379948,379949,379950,379951,379952,379953,379954,379955,379956,379957,379958,379959,379960,379961,379962,379963,379964,379965,379966,379967,379968,379969,379970,379971,379972,379973,379974,379975,379976,379977,379978,379979,379980,379981,379982,379983,379984,379985,379986,379987,379988,379989,379990,379991,379992,379993,379994,379995,379996,379997,379998,379999},
	}
	Banks["us_bankofamerica"] = &Bank{
		"bankofamerica","us","Bank of America","Bank of America","https://www.bankofamerica.com/","#d4001a",
		[]int{374314,374315,374316,374317,374318,374319,374320,374321,374322,374324,374630,374631,374632,374633,374634,374635,374636,374637,374638,374639,374712,374713,374714,374715,374716,374717,374718,374719,374720,374721,377460,377461,377462,377463,377464,377465,377466,377467,377468,377469,400275,400337,400338,400390,400828,401868,401901,401902,401975,401976,401977,402021,402074,402075,402076,402080,402083,402139,402300,402301,402302,402303,402304,402305,402400,402402,402408,402411,402413,402421,402422,402423,402424,402441,402442,402443,402444,402445,402446,402447,402448,402449,402450,402451,402452,402460,402897,403077,403527,403528,403529,403530,403531,403536,403647,403941,403943,405005,405084,405085,405086,405087,405316,405317,405333,405397,405436,405801,405826,406028,406059,407000,407001,407003,407054,407100,407101,407114,407123,407129,407145,407152,407189,407190,407199,407259,407260,408115,408119,408133,408585,409027,409029,409159,409406,409540,409900,410400,410401,410402,410675,410800,410801,410802,410810,410811,410812,410814,410818,410820,410821,410822,410824,411533,411770,411771,411772,411773,411774,411775,411776,411777,411778,411779,411862,412115,412136,412158,412161,412400,412401,412407,412408,412409,412410,412411,412412,412413,412414,412415,412416,412417,412419,412420,412421,412422,412423,412424,412426,412428,412429,412431,412437,412438,412439,412440,412441,412442,412443,412444,412445,412450,412458,412467,412479,412480,412600,412622,413024,413025,413056,413400,413489,413571,413572,413573,413574,413575,413576,413577,413578,413579,413580,414400,414450,414470,414473,414707,414708,414716,414725,414732,414733,414734,414735,414736,414737,414739,414742,414770,414771,414772,414773,415214,415215,415217,415386,415387,415908,415918,415928,415940,416032,416033,416034,416035,416801,416875,416891,417004,417008,417009,417010,417011,417806,417807,417904,419200,419201,419202,419203,419204,419206,419207,419452,419453,419500,419501,419502,419900,419901,420218,420969,420970,421763,421764,421765,421766,421796,422321,422928,422942,423122,423123,423160,423161,423162,423163,423164,423165,423166,423167,423318,423325,423336,423790,424100,424121,424185,424199,424681,425362,425406,425500,425600,425611,425613,425614,425615,425616,425617,425618,425627,425628,425629,425630,425631,425632,425633,425634,425635,425636,425637,425647,425851,426206,426208,426225,426226,426243,426281,426300,426341,426342,426357,426390,426396,426428,426429,426437,426450,426451,426452,426453,426465,426500,426900,427001,427002,427003,427004,427005,427007,427008,427511,427560,428000,428047,428200,428763,429097,429411,429483,429491,429495,429717,429806,430028,430111,430143,430156,430178,430321,430536,430544,430546,430549,430550,430594,430646,430880,430881,430882,430883,430884,431036,431062,431066,431080,431098,431200,431201,431202,431212,431215,431220,431267,431283,431300,431301,431302,431303,431304,431305,431307,431308,431309,431351,431352,431388,431389,431600,431621,431630,431657,431701,431900,431901,431903,431904,432007,432624,432625,432626,432627,432628,432629,432630,432636,432640,432646,432667,432675,432681,432682,432683,432689,432690,432696,432924,433200,433201,433207,433208,433209,433210,433212,433213,433221,433222,433223,433227,433233,433260,433261,433264,433306,433403,433404,433405,433600,433601,433700,433718,433734,433993,434008,434009,434058,434059,434200,434201,434207,434208,434210,434211,434212,434213,434214,434215,434216,434219,434229,434230,434235,434241,434246,434249,434266,434292,434296,435272,435273,435274,435275,435276,435277,435278,435279,435280,435281,435282,435283,435284,435285,435286,435287,435288,435289,435365,435600,435602,435603,435604,435605,435606,435607,435608,435609,435610,435613,435615,435616,435617,435618,435619,435620,435621,435622,435623,435630,435631,435632,435633,435634,435635,435636,435638,435639,435640,435642,435643,435644,435645,435647,435648,435649,435650,435652,435654,435655,435657,435680,435681,435682,435683,435684,435685,435686,435687,435688,435689,435709,436100,436107,436108,436189,436619,436697,436800,436802,436804,438176,438641,438642,438643,438644,438645,438646,438713,438729,438731,438745,438849,438855,438865,438877,438880,438892,440178,440218,440219,440806,440844,440853,440861,440876,440878,440891,440893,440894,440976,441561,441710,442621,442622,442623,442624,442625,442626,442627,442628,442629,442657,442658,442706,442710,442711,442712,442743,442773,442777,442820,442845,442862,442869,442887,442898,442933,443009,443027,443087,443277,443278,444367,444393,444395,445503,445504,445505,445506,445507,445508,445509,445510,445511,445513,445514,446301,446316,446319,446333,446450,446451,446452,446455,446457,446500,446501,446504,446506,446507,446508,446711,446825,447013,447054,447082,447619,447620,447642,447652,447664,447779,448072,448118,448126,448140,448141,448143,448217,448218,448219,448227,448243,448244,448245,448247,448248,448290,448441,448569,448610,448611,448612,448616,448680,448681,448682,448684,448686,448687,448688,448778,448813,448814,448854,448855,449155,449156,449341,449380,449381,449533,449545,449547,449549,449813,449888,450026,450066,450067,450068,451751,452972,453253,454657,454658,454659,455175,455176,455189,455605,455618,455619,460239,460500,460537,460579,460711,460712,461035,461106,461107,461241,461242,461245,461383,461653,461691,461693,461901,461902,461903,461904,461905,461906,461907,461908,461909,461910,461911,461912,461913,461914,461915,461916,461917,461918,461919,461920,462407,462408,462409,462533,462535,462548,462549,462579,462591,463272,463273,463274,463275,463276,463277,463278,463279,463280,463281,463282,463283,463284,463285,463286,463287,463288,463289,463550,463551,463552,463560,463572,463573,463574,463575,463576,463577,463578,463579,463580,463581,463582,463583,463584,463585,463586,463587,463588,463589,463819,463900,463901,463902,463903,464200,464201,464304,464309,464310,464700,464723,464725,464727,464728,464764,464798,464872,464873,464874,464875,464876,464877,464878,464879,464880,464881,464882,464883,464884,464885,464886,464887,464888,464889,464909,465611,465612,465650,466119,466191,467013,467524,467525,467526,467530,467554,467595,468006,468703,470306,470520,470620,470621,470780,471501,471502,471503,471509,471520,471526,471529,471544,471546,471569,471570,471574,471596,472792,472793,473138,473153,473155,473366,474209,474406,474408,474410,474411,474412,474413,474414,474472,474473,474474,474475,474476,474477,474478,474479,474480,474481,474482,474483,474484,474485,474486,474487,474488,474489,474680,474681,474682,474683,474685,474686,474687,475400,475580,475581,476579,477518,478700,478826,479163,479249,479250,479381,479441,479802,479809,479814,479819,479824,479828,479837,479847,479870,480010,480011,480012,480013,480014,480146,480147,480148,480260,480261,480262,480263,480705,480707,480718,48150,48151,48152,48153,48154,48155,48156,48157,481570,48158,481587,48159,481590,481665,481687,481940,482003,482800,482828,483000,484201,484203,484390,484572,484573,484574,484575,484576,484577,484578,484579,484580,484581,484582,484583,484584,484585,484586,484587,484588,484589,484590,484593,484594,484595,484596,484597,484598,485100,485101,485102,485103,485104,485105,485398,485633,486223,486422,487043,487090,488800,488860,488865,488889,488890,488891,488892,488893,488894,488895,489113,489117,489228,489700,493781,502463,504870,510253,510254,510256,514954,516049,518606,520001,520003,520007,522000,522010,522020,522030,522040,522049,522050,522060,522070,522080,522800,522801,522803,522804,522805,522806,522807,522811,522812,522813,522814,522821,522822,522823,522852,522853,522854,523928,523929,524700,525400,525402,526700,527301,527302,528000,528900,528901,528902,528903,528905,530000,530001,530002,530003,530005,530050,530051,530052,530053,530058,530080,530270,530560,530561,530562,530567,530569,530570,530940,530950,530951,531020,531240,531450,531470,531471,531472,531473,531474,531475,531476,531477,531478,531479,531527,531528,531556,532699,532810,532814,532900,532901,532902,532903,532904,532905,532906,532907,532908,532909,532911,532912,533010,533850,533851,533856,534200,534201,534203,534205,534208,534209,534210,534211,534213,534216,534218,534220,534221,534222,534224,534225,534230,534231,534232,534234,534235,534288,535100,535101,535102,535103,535104,535105,535106,535107,535108,535109,539364,539365,540015,540071,540079,540126,540146,540162,540170,540193,540203,540301,540304,540305,540306,540308,540377,540378,540381,540509,540513,540521,540582,540629,540631,540665,540778,540779,540783,540784,540793,540794,540812,540813,540823,540827,540851,540852,540878,540943,541053,541054,541057,541157,541211,541319,541337,541349,541365,541441,541455,541467,541524,541535,541545,541574,541596,541632,541706,541762,541786,541808,541818,541836,541880,541884,541967,542014,542018,542076,542088,542096,542141,542163,542174,542183,542224,542252,542253,542267,542294,542404,542409,542424,542506,542562,542572,542576,542579,542594,542811,542900,543005,543015,543023,543084,543104,543105,543145,543208,543225,543405,543407,543435,543493,543619,543681,543704,543705,543730,543847,543849,543891,544003,544059,544062,544118,544162,544163,544186,544210,544230,544262,544264,544265,544285,544300,544318,544601,544619,544701,544702,544703,544704,544705,544706,544707,544708,544717,544718,544719,544773,544804,544911,544912,544949,544950,545078,545079,545270,545362,545385,545399,545661,545682,545891,545892,546082,546279,546280,546606,546632,546633,546634,546636,546662,546705,546835,546836,546837,546839,546845,546846,546847,546849,547206,547209,547225,547243,547245,547255,547259,547263,547285,547301,547321,547339,547340,547341,547342,547349,547415,547487,547489,547497,547544,547575,547579,547704,548001,548063,548064,548120,548121,548122,548330,548701,548709,548710,548720,548770,548778,548779,548789,549033,549035,549041,549050,549060,549061,549099,549100,549105,549116,549191,549236,549633,550535,551967,551973,551974,552331,552332,554173,554174,554291,554292,554293,554294,554295,554296,556718,556719,556761,556810,556811,556812,556813,556816,556838,556880,556881,556882,556886,556887,556888,556927,556928,558637,558638,558742,558765,558845,558846,558854,558855,558856},
	}
	Banks["us_barclays"] = &Bank{
		"barclays","us","Barclays","Barclays","https://www.banking.barclaysus.com/index.html","#28ace2",
		[]int{414706,432745,432746,432747,432748,439707,439708,469596,486895,486896,514021,514210,514735,514736,514887,514888,514889,514890,514891,514892,515665,517751,517788,523917,524049,540278,541021,545525,545526,546638,547783,552321,552322},
	}
	Banks["us_capitalonebank"] = &Bank{
		"capitalonebank","us","Capital One Bank","Capital One Bank","https://www.capitalone.com/","#004779",
		[]int{400112,400229,400294,400344,400393,400498,401797,401854,402118,402605,402621,403060,403444,403611,403652,403694,405608,405616,410608,410673,411358,411507,412174,414311,414709,415417,415540,415557,415573,415912,415938,419310,430523,430572,430598,438864,441702,450214,450215,477596,477597,479124,480213,486236,489966,490915,493109,493422,494065,514026,517805,518274,518282,518621,518720,518811,520197,528604,529107,529115,529123,529131,529149,529156,529164,529172,529180,529198,530279,530287,530295,530709,530758,539619,539700,539858,543211,545534,545756,546622,546630,547281,552851,552869,554931,557009,557017,557108,557116},
	}
	Banks["us_chase"] = &Bank{
		"chase","us","Chase","Chase","https://www.chase.com/","#0f5ba7",
		[]int{401135,401136,401704,401707,401770,401829,401867,402297,402298,402936,402937,403073,403074,403075,403148,403149,403194,403210,403211,403212,403213,403214,403215,403482,403621,403690,403728,403729,403780,404519,404589,404595,404599,405006,405037,405038,405357,405358,405359,405432,405433,405603,405604,405607,405737,406030,406032,406042,406044,406045,406047,406061,406068,406087,406089,407126,407156,407158,407161,407166,407181,408016,408017,408018,408031,408032,408033,408161,408162,408567,409491,410200,410250,410413,410493,410499,410674,410676,410677,410678,411100,411155,411170,411200,411201,411202,411203,411205,411355,411356,411400,411416,411417,411421,411427,411431,411440,411441,411442,411443,411444,411445,411446,411447,411448,411449,411450,411451,411456,411459,411460,411461,411816,411817,411821,411829,411837,411843,411853,412112,412138,412451,412452,412453,412454,413185,413590,413689,413829,414239,414720,415156,415859,416428,416804,416808,416851,416859,417013,417014,417015,417016,417017,417018,417019,417020,417021,417022,418418,418435,420236,420237,420767,420983,421150,421151,421152,421155,421156,421157,421702,421736,421739,422399,422406,422500,422581,422582,422583,422584,422585,422586,422587,422588,422589,422590,422591,422592,422593,422594,422595,422596,422597,422598,422599,422600,422601,422602,422603,422604,422605,422606,422607,422608,422609,422610,422611,422612,422613,422614,422615,422616,422617,422618,422619,422631,422632,422633,422634,422635,422636,422638,422645,422650,422651,422652,422653,422654,422655,422656,422657,422658,422659,422661,422662,422663,422664,422665,422666,422667,422668,422669,422673,422680,422681,422682,422683,422684,422685,422686,422687,422688,422689,422690,422691,422692,422693,422694,422695,422696,422697,422733,422785,422786,422787,422933,422934,423276,423323,423330,424615,424617,424631,424634,424649,424650,425000,425001,425002,425003,425004,425005,425006,425007,425008,425009,425010,425011,425012,425013,425014,425015,425016,425017,425018,425019,425020,425021,425022,425023,425024,425025,425026,425027,425028,425029,425030,425031,425032,425033,425034,425035,425036,425037,425038,425039,425040,425041,425042,425043,425044,425045,425046,425047,425048,425049,425050,425051,425052,425053,425054,425055,425056,425057,425058,425059,425060,425061,425062,425063,425064,425065,425066,425067,425068,425069,425070,425071,425072,425073,425074,425075,425076,425077,425078,425079,425080,425081,425082,425083,425084,425085,425086,425087,425088,425089,425090,425091,425092,425093,425094,425095,425096,425097,425098,425099,425121,425300,425329,425330,425331,425332,425333,425334,425335,425336,425337,425338,425339,425453,425454,425455,425456,425457,425458,425459,425460,425462,425463,425466,425467,425853,426000,426004,426030,426202,426220,426248,426251,426252,426253,426255,426256,426257,426258,426259,426260,426261,426262,426263,426264,426265,426267,426268,426269,426270,426273,426277,426290,426295,426605,426648,426649,426650,426651,426652,426653,426655,426656,426681,426682,426683,426684,426685,426686,426688,426690,426691,426692,427520,428208,428262,428263,429143,429146,429380,429479,429810,430154,430165,430326,430587,431000,431001,431004,431018,431051,431055,431175,431229,431604,431692,431756,431800,432016,432017,432018,432209,432300,432301,432302,432303,432304,432515,432516,432520,432536,432538,432542,432546,432552,432553,432562,432802,432803,432814,432870,432871,433206,433216,433219,433237,433272,433457,433721,433722,433723,433724,433725,433726,433727,433728,433729,433730,433731,433732,433733,433735,434129,434130,434131,434132,434634,435712,435785,435786,435787,435788,436610,436611,436612,436613,436614,436615,436616,436617,436667,438401,438402,438403,438404,438405,438406,438407,438408,438409,438410,438420,438422,438424,438425,438427,438428,438432,438451,438452,438453,438488,438496,438497,438756,438758,438794,438798,438817,438852,438854,438857,438865,439337,440803,440804,440847,440855,440858,440862,440882,441103,441104,441105,441596,441597,441598,441711,441712,441713,441714,441716,441721,442347,442510,442513,442720,442732,442742,442755,442756,442800,442807,442823,442837,442847,442856,442867,443085,443153,443154,443157,443577,443578,444213,444250,444280,444281,444282,444284,444285,444287,444290,444300,444305,444306,444308,444309,444310,444311,444312,444313,444314,444315,444316,444317,444318,444319,444321,444322,444323,444327,444330,444332,444333,444336,444344,444350,444351,444353,444355,444356,444359,444362,444363,444364,444370,444371,444372,444373,444379,444380,444381,444382,444383,444384,444385,444386,444387,444396,444398,444400,444401,444403,444404,444405,444406,444407,444408,444409,446007,446060,446080,446802,446811,447012,447014,447016,447084,447201,447241,447639,447945,448142,448156,448194,448294,448295,448466,448469,448476,448590,448704,449132,449147,449148,449167,449184,449450,449451,449832,449846,449891,450952,456323,456331,461010,461043,461046,461090,461091,461092,461620,461900,461921,461922,461923,461924,461925,461926,461927,461928,461929,461930,461931,461932,461933,461934,461935,461936,461937,461938,461939,461940,461941,461942,461943,461944,461945,461946,461947,461948,461949,461950,461951,461952,461953,461954,461955,461956,461957,461958,461959,461960,461961,461962,461963,461964,461965,461966,461967,461968,461969,461970,461971,461972,461973,461974,461975,461976,461977,461978,461979,461980,461981,461982,461983,461984,461985,461986,461987,461988,461989,461990,461991,461992,461993,461994,461995,461996,461997,461998,461999,462052,462303,462500,462507,462557,462580,462592,462700,462701,462809,463600,463673,463817,463839,463841,464000,464001,464002,464003,464004,464005,464006,464007,464008,464009,464010,464011,464012,464013,464014,464015,464016,464017,464018,464019,464020,464021,464022,464023,464024,464025,464026,464027,464028,464029,464030,464031,464032,464033,464034,464035,464036,464037,464038,464039,464040,464041,464042,464043,464044,464045,464046,464047,464048,464049,464050,464051,464052,464053,464054,464055,464056,464057,464058,464059,464060,464061,464062,464063,464064,464065,464066,464067,464068,464069,464070,464071,464072,464073,464074,464075,464076,464077,464078,464079,464080,464081,464082,464083,464084,464085,464086,464087,464088,464089,464090,464091,464092,464093,464094,464095,464096,464097,464098,464099,464600,465000,465100,465286,465401,465402,465500,465501,465502,465503,465504,465505,465506,465507,465508,465509,465510,465511,465512,465513,465514,465515,465516,465517,465518,465519,465520,465521,465522,465523,465524,465525,465526,465527,465528,465529,465530,465531,465532,465533,465534,465535,465536,465537,465538,465539,465540,465541,465542,465543,465544,465545,465546,465547,465548,465549,465550,465551,465552,465553,465554,465555,465556,465557,465558,465559,465560,465561,465562,465563,465564,465565,465566,465567,465568,465569,465570,465571,465572,465573,465574,465575,465576,465577,465578,465579,465580,465581,465582,465583,465584,465585,465586,465587,465588,465589,465590,465591,465592,465593,465594,465595,465596,465597,465598,465599,465640,466100,466108,466161,466219,467300,467301,467302,467303,467304,467305,467306,467307,467308,467309,467310,467311,467312,467313,467314,467315,467316,467317,467318,467319,467320,467321,467322,467323,467324,467325,467326,467327,467328,467329,467330,467331,467332,467333,467334,467335,467336,467337,467338,467339,467340,467341,467342,467343,467344,467345,467346,467347,467348,467349,467350,467351,467352,467353,467354,467355,467356,467357,467358,467359,467360,467361,467362,467363,467364,467365,467366,467367,467368,467369,467370,467371,467372,467373,467374,467375,467376,467377,467378,467379,467380,467381,467382,467383,467384,467385,467386,467387,467388,467389,467390,467391,467392,467393,467394,467395,467396,467397,467398,467399,467509,467510,467620,467800,467801,467802,467803,467804,467805,467806,467807,467808,467809,467810,467811,467812,467813,467814,467815,467816,467817,467818,467819,467820,467821,467822,467823,467824,467825,467826,467827,467828,467829,467830,467831,467832,467833,467834,467835,467836,467837,467838,467839,467840,467841,467842,467843,467844,467845,467846,467847,467848,467849,467850,467851,467852,467853,467854,467855,467856,467857,467858,467859,467860,467861,467862,467863,467864,467865,467866,467867,467868,467869,467870,467871,467872,467873,467874,467875,467876,467877,467878,467879,467880,467881,467882,467883,467884,467885,467886,467887,467888,467889,467890,467891,467892,467893,467894,467895,467896,467897,467898,467899,468005,468121,468122,469100,469101,469102,469103,469104,469105,469106,469107,469108,469109,469110,469111,469112,469113,469114,469115,469116,469117,469118,469119,469120,469121,469122,469123,469124,469125,469126,469127,469128,469129,469130,469131,469132,469133,469134,469135,469136,469137,469138,469139,469140,469141,469142,469143,469144,469145,469146,469147,469148,469149,469150,469151,469152,469153,469154,469155,469156,469157,469158,469159,469160,469161,469162,469163,469164,469165,469166,469167,469168,469169,469170,469171,469172,469173,469174,469175,469176,469177,469178,469179,469180,469181,469182,469183,469184,469185,469186,469187,469188,469189,469190,469191,469192,469193,469194,469195,469196,469197,469198,469199,469200,469201,469202,469203,469204,469205,469206,469207,469208,469209,469210,469211,469212,469213,469214,469215,469216,469217,469218,469219,469220,469221,469222,469223,469224,469225,469226,469227,469228,469229,469230,469231,469232,469233,469234,469235,469236,469237,469238,469239,469240,469241,469242,469243,469244,469245,469246,469247,469248,469249,469250,469251,469252,469253,469254,469255,469256,469257,469258,469259,469260,469261,469262,469263,469264,469265,469266,469267,469268,469269,469270,469271,469272,469273,469274,469275,469276,469277,469278,469279,469280,469281,469282,469283,469284,469285,469286,469287,469288,469289,469290,469291,469292,469293,469294,469295,469296,469297,469298,469299,471202,471203,471563,471565,471724,471849,473135,473139,473142,473150,473161,473162,473188,473203,473345,473346,473663,473702,474663,475050,475051,475053,475054,475055,475056,475057,476168,476197,476530,476746,476767,476884,476890,477510,477515,477516,477517,477519,477520,477521,477523,478200,478822,478823,478824,478825,478916,479132,479133,479135,479262,479496,479833,479857,479860,480321,480418,480425,480452,480453,480633,480634,480700,480708,480731,482011,483900,483910,483920,484252,484253,484254,484255,484256,484257,484258,484259,484260,484261,484262,484263,484264,484265,484266,484267,484502,484503,485173,485174,485731,486462,486732,486742,486792,486794,486796,487202,489739,510160,510163,510210,510213,510270,510273,510290,510293,510543,510600,510603,510610,510613,510700,510703,510730,510733,510738,510790,510930,511000,511001,511002,511003,511004,511005,511006,511007,511008,511009,511010,511011,511012,511013,511014,511015,511016,511017,511018,511019,511020,511021,511022,511023,511024,511025,511026,511027,511028,511029,511030,511031,511032,511033,511034,511035,511036,511037,511038,511039,511040,511041,511042,511043,511044,511045,511046,511047,511048,511049,511050,511051,511052,511053,511054,511055,511056,511057,511058,511059,511060,511061,511062,511063,511064,511065,511066,511067,511068,511069,511070,511071,511072,511073,511074,511075,511076,511077,511078,511079,511080,511081,511082,511083,511084,511085,511086,511087,511088,511089,511090,511091,511092,511093,511094,511095,511096,511097,511098,511099,511100,511101,511102,511103,511104,511105,511106,511107,511108,511109,511110,511111,511112,511113,511114,511115,511116,511117,511118,511119,511120,511121,511122,511123,511124,511125,511126,511127,511128,511129,511130,511131,511132,511133,511134,511135,511136,511137,511138,511139,511140,511141,511142,511143,511144,511145,511146,511147,511148,511149,511150,511151,511152,511153,511154,511155,511156,511157,511158,511159,511160,511161,511162,511163,511164,511165,511166,511167,511168,511169,511170,511171,511172,511173,511174,511175,511176,511177,511178,511179,511180,511181,511182,511183,511184,511185,511186,511187,511188,511189,511190,511191,511192,511193,511194,511195,511196,511197,511198,511199,511200,511201,511202,511203,511204,511205,511206,511207,511208,511209,511210,511211,511212,511213,511214,511215,511217,511218,511219,511220,511221,511222,511224,511227,511228,511230,511231,511232,511233,511234,511235,511239,511242,511243,511244,511261,511262,511263,511264,511265,511266,511268,511273,511274,511275,511276,511277,511278,511279,511281,511282,511283,511284,511288,511293,511294,511296,511298,511300,511301,511302,511303,511304,511305,511306,511307,511308,511309,511310,511311,511312,511313,511314,511315,511316,511317,511318,511319,511320,511321,511322,511323,511324,511325,511326,511327,511328,511329,511330,511331,511332,511333,511334,511335,511336,511337,511338,511339,511340,511341,511342,511343,511344,511345,511346,511347,511348,511349,511350,511351,511352,511353,511354,511355,511356,511357,511358,511359,511360,511361,511362,511363,511364,511365,511366,511367,511368,511369,511370,511371,511372,511373,511374,511375,511376,511377,511378,511379,511380,511381,511382,511383,511384,511385,511386,511387,511388,511389,511390,511391,511392,511393,511394,511395,511396,511397,511398,511399,511400,511401,511402,511403,511404,511405,511406,511407,511408,511409,511410,511411,511412,511413,511414,511415,511416,511417,511418,511419,511420,511421,511422,511423,511424,511425,511426,511427,511428,511429,511430,511431,511432,511433,511434,511435,511436,511437,511438,511439,511440,511441,511442,511443,511444,511445,511446,511447,511448,511449,511450,511451,511452,511453,511454,511455,511456,511457,511458,511459,511460,511461,511462,511463,511464,511465,511466,511467,511468,511469,511470,511471,511472,511473,511474,511475,511476,511477,511478,511479,511480,511481,511482,511483,511484,511485,511486,511487,511488,511489,511490,511491,511492,511493,511494,511495,511496,511497,511498,511499,511500,511501,511502,511503,511504,511505,511506,511507,511508,511509,511510,511512,511514,511519,511520,511521,511522,511523,511524,511528,511530,511532,511533,511534,511536,511538,511539,511540,511543,511544,511548,511553,511555,511559,511562,511563,511564,511566,511567,511568,511572,511573,511574,511578,511579,511585,511586,511590,511592,511593,511594,511595,511600,511601,511602,511603,511604,511605,511607,511608,511609,511610,511614,511620,511622,511642,511647,511650,511654,511656,511661,511664,511665,511667,511668,511669,511671,511673,511674,511676,511677,511678,511679,511680,511681,511682,511683,511684,511685,511686,511688,511689,511690,511691,511692,511693,511694,511695,511696,511698,511699,511712,511716,511721,511726,511730,511732,511733,511734,511740,511742,511743,511744,511746,511747,511748,511749,511756,511760,511762,511763,511765,511766,511767,511770,511771,511774,511775,511778,511779,511780,511781,511782,511783,511784,511785,511788,511790,511791,511793,511794,511796,511799,511800,511807,511808,511809,511812,511813,511814,511815,511816,511817,511818,511819,511820,511821,511822,511823,511824,511825,511826,511827,511828,511829,511830,511831,511832,511833,511834,511836,511837,511838,511839,511840,511841,511842,511843,511844,511847,511848,511849,511850,511851,511852,511853,511856,511857,511859,511860,511861,511864,511865,511866,511867,511869,511870,511871,511872,511874,511875,511876,511878,511879,511880,511881,511882,511885,511887,511888,511889,511890,511891,511892,511893,511894,511895,511896,511897,511898,511899,511900,511901,511902,511904,511910,511911,511913,511916,511917,511920,511921,511932,511934,511937,511938,511939,511942,511953,511958,511964,511967,511972,511974,511981,511987,511988,511995,511998,511999,512012,512064,512065,512257,514018,514025,514027,514058,514182,514193,514230,514377,514427,514504,514543,514573,514873,514874,514901,514909,514918,514922,514923,514925,514938,515563,515908,517945,517954,517955,517958,517966,517975,517976,517993,518015,518016,518017,518021,518022,518023,518026,518027,518041,518047,518052,518053,518055,518059,518065,518069,518199,518243,518245,518316,518336,518337,518338,518445,518448,518449,518453,518495,518698,518748,518753,518818,518863,520064,520065,520353,520364,521115,521116,521127,521128,521142,521146,521147,521148,521150,521151,521163,521181,521843,521845,522276,522277,522278,522297,523001,523002,523003,523010,523028,523030,523080,523800,523801,523809,524002,524003,524017,524018,524020,524026,524027,524034,524100,526010,526020,526021,526022,526023,526024,526025,526028,526031,526032,526033,526034,526035,526036,526037,526044,526045,526049,526050,526053,526058,526060,526065,526066,526072,526077,526078,526080,526082,526084,526087,526090,526300,526307,526315,526316,526317,526327,526338,526340,526341,526343,526344,526345,526346,526347,526348,526349,526350,526353,526378,526379,526386,526388,526394,528600,528601,528602,528603,528604,528605,528606,528607,528608,528609,528610,528611,528612,528613,528614,528615,528616,528617,528618,528619,528620,528621,528622,528623,528624,528625,528626,528627,528628,528629,528630,528631,528632,528633,528634,528635,528636,528637,528638,528639,528640,528641,528642,528643,528644,528645,528646,528647,528648,528649,528650,528651,528652,528653,528654,528655,528656,528657,528658,528659,528660,528661,528662,528663,528664,528665,528666,528667,528668,528669,528670,528671,528672,528673,528674,528675,528676,528677,528678,528679,528680,528681,528682,528683,528684,528685,528686,528687,528688,528689,528690,528691,528692,528693,528694,528695,528696,528697,528698,528699,528707,528790,529800,529801,529802,530021,530022,530023,530024,530027,530028,530029,530250,530251,530252,530253,530254,530255,530256,530257,530258,530259,531570,531571,531572,531573,531574,531575,531576,531577,531578,531579,531600,531601,531602,531603,531604,531605,531606,531607,531608,531609,532300,532350,532351,532352,532353,532354,532355,532356,532357,532358,532359,532521,532590,534710,534711,534712,534713,534714,534715,534716,534717,534718,534719,534810,534811,534812,534813,534814,534815,534816,534817,534818,534819,536250,536990,536991,536992,536993,536994,539618,540035,540168,540272,540283,540325,540326,540405,540501,540512,540514,540525,540534,540539,540540,540555,540562,540568,540575,540730,540798,540810,540811,540831,540848,540875,541011,541102,541130,541135,541146,541193,541205,541273,541284,541316,541347,541357,541389,541390,541391,541435,541461,541489,541503,541507,541515,541527,541581,541594,541622,541649,541711,541712,541714,541716,541739,541740,541758,541760,541769,541782,541815,541974,542003,542023,542074,542103,542128,542188,542205,542228,542240,542241,542242,542243,542245,542270,542287,542299,542420,542438,542539,542585,542596,542709,542759,542807,542979,543056,543063,543095,543119,543138,543143,543282,543286,543312,543411,543505,543555,543562,543583,543589,543590,543593,543596,543598,543612,543647,543695,543710,543748,543762,543880,544018,544181,544215,544216,544376,544377,544378,544379,544671,544745,544746,544747,544748,544749,544750,545060,545090,545284,545334,545335,545336,545337,545338,545339,545340,545341,545342,545491,545572,545667,545793,545890,545921,546029,546469,546500,546506,546507,546536,546544,546598,546599,546604,546615,546620,546621,546626,546643,546644,546647,546651,546655,546656,546657,546663,546664,546667,546671,546672,546675,546677,546679,546691,546710,546711,546725,546731,546751,547163,547169,547209,547217,547221,547230,547272,547277,547312,547314,547315,547363,547364,547387,547437,547446,547507,547519,547586,547688,547739,547798,547799,547891,547899,548010,548018,548047,548099,548100,548310,548311,548317,548325,548331,548341,548342,548343,548344,548376,548381,548382,548392,548395,549085,549086,549087,549088,549089,549091,549092,549093,549094,549104,549144,549145,549171,549175,549967,550040,550306,550307,550308,550311,550366,551027,552475,552631,553772,554186,554303,554306,554307,554453,554454,554455,554456,554457,554458,554564,554570,554608,554670,554750,554751,554752,554753,554754,554807,554896,554897,554898,554962,556706,556708,556717,556723,556724,556746,556753,556760,556820,556821,556822,556823,556824,556825,556826,556827,556828,556829,556890,556891,556892,556893,556894,556895,556896,556897,556898,556899,556917,558016,558017,558018,558019,558089,558090,558112,558113,558114,558115,558158,558250,558692,558729,558757,558767,558821,558823,565227,567223,568030,588877,588896,589563,589957,601340,603028},
	}
	Banks["us_citibank"] = &Bank{
		"citibank","us","Citibank","Citibank","https://online.citi.com/US/","#0088cf",
		[]int{402835,402836,402856,402889,409216,409217,409218,412926,413777,415763,421198,423353,425727,426534,433388,433389,440669,443289,444700,444780,444784,445167,445450,445984,446270,446325,446545,446555,446556,446863,447742,447963,447964,447965,448400,448448,448488,448502,448600,448601,448602,448603,448604,448605,448606,448607,448608,448609,448670,448671,448672,448673,448674,448675,448676,448677,448678,448679,448700,448701,448702,448703,448721,448722,448734,448735,448736,448741,448742,449227,449228,449441,449800,449877,450900,451449,451450,453210,453248,453971,453972,454178,454179,454325,454413,454415,454416,454605,454750,455038,455390,455596,455643,455644,455716,455987,455988,456001,456313,456318,456365,456366,456397,456452,456487,456547,456548,456738,456744,456822,456873,456881,458102,459801,461657,461669,462120,462121,462122,463669,463670,466507,466516,467092,468524,468525,469548,469549,469551,469552,471580,471586,471588,471597,473113,473116,477795,477943,477944,478300,478306,478355,478356,478358,478360,478380,478400,478405,478406,478455,478456,478457,478460,478480,478490,479110,479111,479149,479150,479340,485738,485801,489962,490140,490696,490698,490902,490930,490931,490932,490933,491002,491832,491833,492053,492075,492082,492087,492200,493137,493138,493139,493714,494122,494123,496646,498500,510296,512068,512103,512106,512107,512108,514012,514034,514035,514054,514533,514571,517700,517762,517940,517949,517950,517969,517983,518371,518390,518391,518429,518607,518752,520100,521800,521801,521802,521803,521804,521805,521806,521807,521808,521809,522258,522281,525600,525601,525602,525603,525604,525605,525606,525607,525608,525609,525610,525650,526218,526219,526220,526221,526222,526223,526224,526225,526226,526227,526228,526229,530878,530890,530891,530892,530893,530894,530895,530896,530897,530898,530899,530900,530901,530902,530903,530904,530905,530906,530907,530908,530909,531160,536218,536219,536220,536221,536222,536223,536224,536225,536226,536227,536228,536229,539600,539603,539610,539611,539620,539640,539641,539642,539643,539644,539646,539647,539655,539656,539680,539681,539682,539810,539813,539820,539821,539822,539840,539841,539842,539843,539844,539846,539855,539856,539857,539860,539870,539871,539880,539882,539883,540165,540175,540184,540302,540432,540541,540553,540554,540598,540621,540680,540751,540828,540855,540924,540938,540975,540999,541065,541242,541299,541304,541332,541484,541505,541526,541566,541572,541610,541632,541655,541692,541693,541694,541696,541705,541757,541759,541862,541887,541931,541955,541961,541972,542019,542061,542064,542092,542177,542379,542415,542418,542471,542472,542531,542556,542617,542655,542785,542831,542947,543008,543021,543060,543185,543219,543298,543449,543594,543666,543667,543780,544129,544282,544283,544293,544339,544700,544979,545137,545715,545716,545987,546002,546068,546294,546527,546613,546616,546619,546627,546628,546646,546648,546653,546695,547147,547148,547175,547200,547218,547297,547431,547439,547641,547786,547787,548036,548068,548161,548399,548501,548503,548647,549102,549113,549114,549137,549143,549146,549149,549152,549154,550594,551070,552060,552081,552093,552815,552830,552832,552841,554305,554426,554577,554619,554637,554863,554902,556043,556709,556800,556801,556802,556803,556804,556805,556806,556807,556808,556809,556870,556871,556872,556873,556874,556875,556876,556877,556878,556879,556900,556921,556924,557127,558101,558800,558802,558828,558832,558837,558857,558858,558859,558861,558862,558863,558879,558880},
	}
	Banks["us_fifththird"] = &Bank{
		"fifththird","us","Fifth Third Bank","Fifth Third Bank","https://www.53.com/","#254094",
		[]int{415786,444512,542432,548009},
	}
	Banks["us_googlewallet"] = &Bank{
		"googlewallet","us","Google Wallet","Google Wallet","https://wallet.google.com/","#4285f4",
		[]int{524825},
	}
	Banks["us_hsbc"] = &Bank{
		"hsbc","us","HSBC","HSBC","https://www.us.hsbc.com/","#dd1f26",
		[]int{401749,403000,403001,403002,403010,403012,403013,403200,403201,403202,403299,403414,403415,404203,404614,404615,404649,404693,404993,406096,407105,408120,408121,412128,412179,412201,412202,412203,412204,412205,412207,412208,412209,412210,412211,412215,412224,412225,412226,412244,412260,412283,413113,413120,413140,413160,413165,413170,416841,420042,420043,420987,422983,422984,423311,423322,424642,424735,425400,425418,425451,425452,425477,425493,426930,426932,426933,426934,426944,426993,426994,426995,426996,427501,427502,427546,429113,429137,429152,429184,429185,429350,430200,430521,431047,431467,431468,431663,431689,431734,434351,434385,438826,438845,438856,438893,438894,440424,440560,440561,442770,442874,443122,443508,446810,447717,447718,449867,449882,449883,452500,452520,452526,463000,463021,463099,463814,463815,464444,464445,466304,466800,467528,471206,473068,474672,474673,474700,475551,477767,478827,478828,478829,479365,480306,480385,480397,491099,491921,493851,498896,512025,512026,512027,512300,512731,514953,515592,515593,515597,515598,515599,515625,517669,517993,521507,521531,522944,525478,526835,531480,532561,532830,532831,532839,534248,539662,540159,540419,540420,540421,540422,540423,540424,540425,540426,540580,540633,540704,540707,540789,540790,540791,540792,540801,541058,541336,541402,541444,541448,541449,541458,541474,541513,541536,541549,541578,541586,541627,541829,542036,542135,542142,542376,542442,542583,542775,543140,543234,543235,543328,543329,543339,543368,543675,543700,543701,543702,543703,543857,543961,544045,544200,544269,544316,544317,544347,544361,544368,544515,544516,544910,545694,545695,545800,546281,546282,546283,546608,546609,546610,546611,546612,546641,546701,546702,547478,548042,548043,548897,548950,548951,548955,549062,549066,549107,549109,549110,549111,549118,549119,549120,549121,549131,549132,549133,549134,549135,549409,549943,549944,549945,550300,552030,552234,552318,552319,554655},
	}
	Banks["us_jpmorganchase"] = &Bank{
		"jpmorganchase","us","JPMorgan Chase","JPMorgan Chase","https://www.jpmorganchase.com/","#ececeb",
		[]int{438857,556708},
	}
	Banks["us_metabank"] = &Bank{
		"metabank","us","Metabank","Metabank","https://www.metabank.com/","#004990",
		[]int{404195,419672,432732,435899,441709,446106,460005,472776,475034,475423,475427,487211,531108},
	}
	Banks["us_navyfederal"] = &Bank{
		"navyfederal","us","Navy Federal Credit Union","Navy Federal Credit Union","https://www.navyfederal.org","#0f4471",
		[]int{400022,403216,406041,406095,406315},
	}
	Banks["us_payoneer"] = &Bank{
		"payoneer","us","Payoneer","Payoneer","https://www.payoneer.com/","#ff4800",
		[]int{530072},
	}
	Banks["us_simple"] = &Bank{
		"simple","us","Simple","Simple","https://www.simple.com","#007cd9",
		[]int{481171},
	}
	Banks["us_suntrust"] = &Bank{
		"suntrust","us","SunTrust","SunTrust","https://www.suntrust.com","#ef7622",
		[]int{442505,451805,486560,519667,519669,546540,552393,553693,553694,557621},
	}
	Banks["us_tdbank"] = &Bank{
		"tdbank","us","TD Bank","TD Bank","https://www.td.com/","#35b233",
		[]int{401777,402944,408537,408540,415874,417207,417208,429482,432855,438766,449163,465652,479213},
	}
	Banks["us_usbank"] = &Bank{
		"usbank","us","US Bank","US Bank","https://www.usbank.com/index.html","#0c2074",
		[]int{404645,414780,420719,424604,436618,446053,470758,476153,521679},
	}
	Banks["us_wellsfargo"] = &Bank{
		"wellsfargo","us","Wells Fargo","Wells Fargo","https://www.wellsfargo.com/","#a00000",
		[]int{401158,401645,401655,401670,401828,402238,402240,402241,403120,403121,403122,403123,403124,403125,403126,403127,403128,403129,403130,403131,403133,403135,403136,403137,403138,403139,403140,403458,403539,403562,403700,403718,403719,403740,403790,404000,404001,404002,404003,404004,404005,404006,404007,404008,404009,404010,404011,404012,404013,404014,404015,404016,404017,404018,404019,404020,404021,404022,404023,404024,404025,404026,404027,404028,404029,404030,404031,404032,404033,404034,404035,404036,404037,404038,404039,404040,404041,404042,404043,404044,404045,404046,404047,404048,404049,404050,404051,404052,404053,404054,404055,404056,404057,404058,404059,404060,404061,404062,404063,404064,404065,404066,404067,404068,404069,404070,404071,404072,404073,404074,404075,404076,404077,404078,404079,404080,404081,404082,404083,404084,404085,404086,404087,404088,404089,404090,404091,404092,404093,404094,404095,404096,404097,404098,404099,404681,405000,405010,405013,405014,405015,405017,405018,405020,405023,405024,405025,405035,405061,405062,405063,405072,405711,405713,405714,406098,407110,407163,407184,407188,407195,408000,408189,409251,409280,409281,409282,409283,409284,409285,409286,409287,409288,409290,409292,409293,409294,409295,409296,409297,409298,409299,410100,411277,411339,411346,411876,412001,412002,412003,412004,412020,412080,412081,412082,412083,412085,412086,412087,412088,412089,412090,412091,412092,412093,412094,412095,412096,412098,412099,412493,412737,412738,413041,413751,413923,414718,414922,414923,414924,414927,414928,414929,414930,414931,414980,414984,414987,414988,414990,414991,414993,414995,414996,414997,414998,414999,415080,415081,415082,415083,415085,415086,415087,415088,415089,415090,415091,415092,415093,415094,415095,415096,415097,415098,415320,415350,415351,415352,415354,415357,415358,415393,415533,415728,415746,415747,416000,416001,416002,416003,416004,416006,416007,416724,416725,416886,418510,418512,418513,418514,418515,418516,418517,418518,418519,418520,418521,418527,418528,418529,420078,420094,420095,420303,420307,420500,420501,420682,421733,422384,422770,422841,422842,422844,422847,422848,422849,422850,422851,422853,422855,422858,422859,422917,422918,422939,422950,422965,423731,424610,424690,424713,424720,424780,424782,424784,424787,424788,424790,424791,424792,424793,424794,424795,424796,424797,424798,424799,425424,425431,426238,426249,426294,426473,426536,426537,426538,426539,426540,426541,426542,426543,426544,426546,426547,426548,426549,426550,426551,426552,426553,426554,426555,426617,426640,426641,426644,426657,426660,426910,426911,426950,426952,426956,426957,426958,426959,426960,426961,426962,426963,426965,426967,426968,427076,427180,427199,428027,428032,429178,429412,429477,429676,430070,430071,430072,430073,430074,430079,430117,430716,430728,430741,430763,430943,431056,431578,431609,431625,431643,431652,431706,431739,431741,432101,432370,432371,432372,432373,432374,432375,432376,432377,432378,432379,432380,432381,432382,432383,432384,432385,432386,432387,432388,432389,432390,432391,432392,432393,432732,432734,432749,432854,433751,433752,433753,433754,433755,433756,433757,433758,433759,433760,433761,433762,433763,433764,433765,433767,433768,433769,433789,435101,435102,435103,435104,435107,435108,435109,435110,435111,435112,435113,435114,435115,435116,435117,435118,435119,435157,435383,435510,435549,435662,435667,435670,436161,436162,436163,436164,436165,436166,436167,436168,436169,436171,436172,436173,436174,436175,436176,436177,436178,436179,436181,436182,437811,438044,438461,438462,438463,438464,438467,438469,438470,438471,438472,438473,438474,438475,438476,438480,438481,438482,438483,438486,438489,438490,438491,438492,438493,438494,438499,438501,438502,438504,438506,438507,438514,438515,438516,438517,438518,438519,438520,438562,438563,438565,438566,438567,438568,438569,438570,438572,438573,438574,438830,438862,438884,438945,438976,440334,440379,441625,441946,442200,442202,442329,442518,442550,442551,442552,442553,442554,442555,442556,442557,442558,442559,442560,442561,442562,442563,442565,442566,442567,442568,442569,442575,442644,442661,442662,442663,442665,442667,442669,442808,443107,443148,443149,443161,443243,443526,443544,443699,445227,446023,446024,446082,446083,446106,446533,446539,446540,446541,446542,446543,447083,447166,447430,447431,447432,447433,447434,447435,447459,447466,447490,447670,447672,447673,447674,447676,447677,447678,447679,447680,447681,447682,447683,447684,447685,447687,447688,447689,447690,447691,447866,447867,447901,447902,447903,447990,448067,448069,448129,448461,448489,448777,448818,448872,449018,449117,449215,449235,449252,449256,449299,449372,449373,449376,449398,449496,449585,449705,449706,449707,449708,449709,449710,449711,449712,449714,449715,449716,449717,449718,449719,449720,449721,449722,449723,450208,450238,450239,450240,450241,450242,450243,450244,450245,450246,450247,450249,450394,450395,450397,450398,450399,453234,453235,453270,453271,453274,453275,453276,453277,453278,453279,453280,453281,453282,453283,453284,453285,453287,453288,453289,453512,453513,453514,453515,453516,453517,453518,453520,453521,453522,453523,453524,453526,453528,453529,453530,453531,453532,453533,453543,453552,455401,460003,460004,460005,460006,460007,460010,460011,460794,461051,461634,462100,462564,462586,463695,463696,463697,463698,463699,463700,463701,463702,463703,463704,463705,463706,463707,463708,463709,463710,463711,463712,463713,463714,464032,466001,466002,466003,466004,466005,466006,466007,466008,466009,466010,466011,466012,466013,466014,466015,466016,466017,466018,466019,466020,466131,467590,467594,467600,467601,467602,467603,467604,467605,467606,467607,467608,467609,467610,467611,467612,467613,467614,467616,467617,467618,467619,469400,469404,469409,469501,469502,469503,469504,469505,469506,469507,469508,469509,469510,469511,469513,469514,469516,469517,469518,469519,469520,470057,470500,470715,470716,470721,470722,470723,470724,470725,470727,470728,470729,470731,470732,470733,470734,470735,470736,470737,470738,470739,470740,472099,472101,472102,472103,472104,472105,472106,472107,472108,472160,472162,472164,472165,472400,472600,472601,472602,472603,472604,472605,472606,472607,472608,472609,472610,472611,472612,472613,472614,472615,472616,472617,472618,472619,472620,472621,472622,472623,472624,472625,472626,472627,472628,472629,472630,472631,472632,472633,472634,472635,472636,472637,472638,472639,472640,472641,472642,472643,472644,472645,472646,472647,472648,472649,472650,472651,472652,472653,472654,472655,472656,472657,472658,472659,472660,472661,472662,472663,472664,472665,472666,472667,472668,472669,472670,472671,472672,472673,472674,472675,472676,472677,472678,472679,472680,472681,472682,472683,472684,472685,472686,472687,472688,472689,472690,472691,472692,472693,472694,472695,472696,472697,472698,472699,472721,472734,473099,473100,473112,473140,473168,473169,473211,473378,473379,473381,473385,473386,473391,473392,473393,473394,473395,474100,474300,474302,474305,475100,475101,475102,475103,475104,475105,475106,475107,475108,475540,475585,475602,475603,475604,475608,475609,475637,475675,475676,475683,476163,476300,476301,476302,476304,476305,476306,476307,476310,476312,476313,477306,477307,477308,477310,477311,477312,477313,477315,477316,477318,477319,477320,477321,477322,477324,477325,477326,477327,477328,477329,477330,477331,477332,477333,477334,477335,477336,477337,477338,477340,477341,477726,477738,477739,477880,477881,477882,477885,477887,477888,477894,477895,477896,477899,478844,478861,478890,478892,478894,478896,478897,478898,478899,478963,478966,478967,478971,478973,478975,478976,478977,478978,478979,478980,478981,478982,478983,479346,479810,480356,480358,480359,480361,480362,480363,480365,480366,480367,480368,480369,480372,480374,480396,480801,481942,482100,482900,484400,484405,484500,484504,485200,485620,486010,486101,486102,486103,486104,486105,486106,486107,486108,486109,486110,486111,486113,486114,486115,486117,486118,486119,486120,486121,486122,486123,486124,486125,486127,486128,486129,486130,486131,486132,486133,486135,486136,486137,486138,486139,486140,486827,486830,486831,486832,486848,486850,486851,487309,487322,487360,487451,491577,491980,491981,491982,491983,491984,491985,491986,491987,491988,491989,491990,491991,491992,491993,514019,514020,514120,514121,514122,514123,514124,514125,514126,514127,514128,514129,517834,517844,517854,527700,527701,527702,527703,527704,527705,527706,527707,527708,527709,527710,527711,527712,527713,527714,527715,527716,527717,527718,527719,527720,527721,527722,527723,527724,527725,527726,527727,527728,527729,527730,527731,527732,527733,527734,527735,527736,527737,527738,527739,527740,527741,527742,527743,527744,527745,527746,527747,527748,527749,527750,527751,527752,527753,527754,527755,527756,527757,527758,527759,527760,527761,527762,527763,527764,527765,527766,527767,527768,527769,527770,527771,527772,527773,527774,527775,527776,527777,527778,527779,527780,527781,527782,527783,527784,527785,527786,527787,527788,527789,527790,527791,527792,527793,527794,527795,527796,527797,527798,527799,528200,528201,528202,528203,528204,528205,528206,528207,528208,528209,528230,528231,528235,530070,530074,530079,531700,531701,532580,532584,532589,533150,533151,533152,533153,533154,533155,533156,533157,533158,533159,533470,533471,533472,540016,540177,540510,540597,541018,541037,541372,541424,541426,541447,541570,541571,541620,541722,541821,541828,541830,541831,541832,541834,541837,541902,542057,542143,542567,542586,542950,543693,543825,544436,544735,545326,545404,545885,546782,547224,547240,547366,547463,547464,547547,547643,547802,548077,549084,549095,549096,549097,549098,549103,549140,549170,549714,549715,550315,556905,556919,556931,556939,558370,558517,558618,558668,558669,558674,558835},
	}
	Banks["uy_brou"] = &Bank{
		"brou","uy","Banco de la República Oriental del Uruguay","Bank of the Oriental Republic of Uruguay","http://www.brou.com.uy/","#024181",
		[]int{421821,494016},
	}
	Banks["uy_santander"] = &Bank{
		"santander","uy","Santander","Santander","http://www.santander.com.uy/","#ff0000",
		[]int{534263,545703},
	}
}