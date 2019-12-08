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
- Outputs links based on hardcoded initial single index list starting URL

``` bash
> go run eroscrape.go
ExperienceVaultURL:  https://erowid.org/experiences/exp.cgi?S1=15&ShowViews=0&Cellar=0&Start=0&Max=500
ExperienceVaultURL:  https://erowid.org/experiences/exp.php?ID=16996
ExperienceVaultURL:  https://erowid.org/experiences/exp.php?ID=16973
ExperienceVaultURL:  https://erowid.org/experiences/exp.php?ID=29874
ExperienceVaultURL:  https://erowid.org/experiences/exp.php?ID=11686
ExperienceVaultURL:  https://erowid.org/experiences/exp.php?ID=11218
ExperienceVaultURL:  https://erowid.org/experiences/exp.php?ID=3286
ExperienceVaultURL:  https://erowid.org/experiences/exp.php?ID=92059
ExperienceVaultURL:  https://erowid.org/experiences/exp.php?ID=94567
ExperienceVaultURL:  https://erowid.org/experiences/exp.php?ID=67153
ExperienceVaultURL:  https://erowid.org/experiences/exp.php?ID=71409
ExperienceVaultURL:  https://erowid.org/experiences/exp.php?ID=78740
ExperienceVaultURL:  https://erowid.org/experiences/exp.php?ID=73483
```
...
``` bash
DaturaExperienceURL:  https://erowid.org/experiences/exp.php?ID=29
DaturaExperienceURL:  https://erowid.org/experiences/exp.php?ID=28
```

_**Next Up:**_
- Fetching the experience report text from each URL gathered.
- Parse the report into plain text die to the vault's often variant consistancy on text encoding mistakes and report page HTML layout
