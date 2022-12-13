import fs from "fs";
import https from "https";

async function getImg() {
    let num = Math.round(Math.random() * 2);
    if (!num || num == 1) {
        getCatBoys();
        return;
    }
    getNekos();
}

async function getCatBoys() {
    https.get("https://api.catboys.com/img ", (res) => {
        console.log('from catboys')
        res.on("data", async (data) => {
            var {
                url
            } = JSON.parse(data.toString())
            https.get(url, cum => {
                cum.on('data',cuM=>{
                    console.log(cuM) 
                })
            })
        });
    });
};

async function getNekos() {
    https.get("https://nekos.best/api/v2/neko", (res) => {
        console.log('from nekoBoys')
        res.on("data", async (data) => {
            let {
                url
            } = JSON.parse(data.toString()).results[0]
            https.get(url, cum => {
                cum.on('data',cuM=>{
                    console.log(cuM) 
                })
            })
        });
    });
}
getImg();
