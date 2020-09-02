你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
# erowid-a-scrape
### Golang based [Erowid Experience Vaults](https://erowid.org/experiences/exp_front.shtml) scraper helper code made with [Colly](https://github.com/gocolly/colly)

---
**Motivation**

Giving an [older project](https://github.com/Bedrovelsen/erowid-a-scrape/raw/master/daturagen3.png) relating to datura experince reports new life with the goals being

>>>
1) Automating text corpus creation instead of former project's manually fetched & parsed thus error filled text corpus file
2) Applying the text corpus generation and prior projects single substance only to any substance with reports in the erowid experince vaults via dropdown picker just like on the home page of erowids experience vaults.

---

_**Current Functionality**:_
- Grabs report links based on hardcoded vault category index page URL and outputs to links.lst
- Appends report text from each page in list of links generated to master corpusText and outputs to corpus.text

``` bash
> go run eroscrape.go
Starting erowid-a-scrape
[erowid experience report text corpus generation]

Report index base URL: https://erowid.org/experiences/exp.cgi?S1=15&ShowViews=0&Cellar=0&Start=0&Max=500
Finished getting report URLs
Report URLs written to file links.lst
Finished fetching reports into text corpus
Text corpus written to file corpus.txt
Erowid-a-scrape finished

> tail links.lst 
https://erowid.org/experiences/exp.php?ID=2862
https://erowid.org/experiences/exp.php?ID=2819
https://erowid.org/experiences/exp.php?ID=1829
https://erowid.org/experiences/exp.php?ID=1828
https://erowid.org/experiences/exp.php?ID=1827
https://erowid.org/experiences/exp.php?ID=1821
https://erowid.org/experiences/exp.php?ID=1820
https://erowid.org/experiences/exp.php?ID=1819
https://erowid.org/experiences/exp.php?ID=29
https://erowid.org/experiences/exp.php?ID=28

 I walked around the house talked to people even though I was making people do things by thought, until I realized I was still in bed.
 I couldn't come to any conclusion.
 I slept for maybe six hours and with great difficulty got out of bed.
 However, once out everything was clear, and normal.
 No pupil dilation, some dry throat, but otherwise normal.
 If not a bit relaxed and clear headed.
 For me it was definitely a pleasurable experience.
```

_**Next Up:**_
- ~~Fetching the experience report text from each URL gathered.~~ [Done]
- Parse / clean up the corpus text into plain formatting type due to he vault's often variant consistancy on text encoding mistakes and report page HTML layout
