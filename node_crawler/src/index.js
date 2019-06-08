const fetch = require("node-fetch")
const curl =`
curl 'http://www.wrd.cn/view/home/bignessEvent/getBignessEventList.action' 
-H 'Origin: http://www.wrd.cn' 
-H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36' 
-H 'Content-Type: application/x-www-form-urlencoded; charset=UTF-8' 
-H 'Accept: */*' -H 'Referer: http://www.wrd.cn/home.shtml?page=2' 
-H 'X-Requested-With: XMLHttpRequest' 
-H 'browser-w: 1440' -H 'browser-h: 740' 
--data 'startTime=2019-06-02+05%3A30%3A15&endTime=2019-06-09+05%3A30%3A15' 
--compressed`
fetch("http://www.wrd.cn/view/home/bignessEvent/getBignessEventList.action",{
    methods: "POST",
    body:'startTime=2019-06-02+05%3A30%3A15&endTime=2019-06-09+05%3A30%3A15',
    headers:{
        "Origin": "http://www.wrd.cn",
        "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36",
        "Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
        "Accept":"*/*",
        "Referer": "http://www.wrd.cn/home.shtml?page=2",
        "X-Requested-With": "XMLHttpRequest"
    }
}).then(res =>{
    console.log(res.json)
})