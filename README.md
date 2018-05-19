# 全球银行卡号码校验与信息查询

## 相关资料
* [alipay](https://ccdcapi.alipay.com/validateAndCacheCardInfo.json?cardNo=6222005865412565805&cardBinCheck=true)银行卡号接口
* [Payment card number](https://en.wikipedia.org/wiki/Payment_card_number)

## 银行卡号校验

[Luhn检验数字算法](https://en.wikipedia.org/wiki/Luhn_algorithm)（Luhn Check Digit Algorithm），也叫做模数10公式，是一种简单的算法，用于验证银行卡、信用卡号码的有效性算法。对所有大型信用卡公司发行的信用卡都起作用，这些公司包括美国Express、护照、万事达卡、Discover和用餐俱乐部等。这种算法最初是在20世纪60年代由一组数学家制定，现在Luhn检验数字算法属于大众，任何人都可以使用它。

算法：将每个奇数加倍和使它变为单个的数字，如果必要的话通过减去9和在每个偶数上加上这些值。如果此卡有效，那么结果必须是10的倍数。
