// ==UserScript==
// @name         JetBrains Activation Code Generator
// @namespace    https://github.com/
// @version      1.0
// @license MIT
// @description  Add a button on the plugin homepage and click to get the plugin activation code
// @author       tunknown
// @Credits      https://greasyfork.org/en/users/1335977-weiba007
// @match        https://plugins.jetbrains.com/*
// @grant        GM.setClipboard
// @grant        GM.addStyle
// @grant        window.onurlchange
// @downloadURL https://update.greasyfork.org/scripts/525226/JetBrains%20Activation%20Code%20Generator.user.js
// @updateURL https://update.greasyfork.org/scripts/525226/JetBrains%20Activation%20Code%20Generator.meta.js
// ==/UserScript==

const pemEncodedKey = `-----BEGIN PRIVATE KEY-----
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
sm0eUeblKh2/BExSaZrJ4WV10lDZWw==
-----END PRIVATE KEY-----`;


const pemEncodedCrt = `-----BEGIN CERTIFICATE-----
MIIE2TCCAsGgAwIBAgIRAN7BFfgl0UuGNAbuts/u2c0wDQYJKoZIhvcNAQELBQAw
GDEWMBQGA1UEAxMNSmV0UHJvZmlsZSBDQTAeFw0yNTA2MTQwNjA5NDFaFw0zNTA2
MTUwNjA5NDFaMBYxFDASBgNVBAMTC2pldGJyYS1mcmVlMIICIjANBgkqhkiG9w0B
AQEFAAOCAg8AMIICCgKCAgEA1Rl8Hz2vS4rXfm7dYJJ5qcE6Y72KNXDFbCTuwqxZ
8AXBMT/hfkF1RhUvnJzn/gbDqjkTkOfZyebjGSi7Hf9B6ns566US+5fPdkWuZPo/
jdxqWk88vkOGFG1ZQjgbxMZkLnTm/llJFcn4b3OhGROPamSL9+G+2fXsfvAR0DS5
4P5IiBMzREJ12AJ6ibNUYlaA1JtT4bfgo0W/vi5kFR1eKJlYPoax6q3vmpQn7mb+
5VkzsvIOZnnejubFMkG82ZA/ethQ4AWqQkd0hUI1kiAhct71apa2joGil6em6uFd
WsOrCfcjGKCjB5hZHla/lsec6m5i8YHZP5PDKF31U1Ar6flVNR1auVlGkM7PW/Eu
S2+KHv3QEuWd/TllzNQfU1kfeRCetvendtdVpbg57woZrW4A0o+kIYCAu+Lr7NxJ
gwnFcXeyP8NyHXyQVUJ3hOIjHA/9+b8+NCyRe+/RtyZm8cD824439YkpUnXWseaJ
mGQcPNE/QZ5dJVK+HAQf3dnpxtmnqotIaViKk+ynWdxUz+/RDWBabXUaT9VKh81c
GZPLmcoBcfKiWmFnCBTzqiorjC8MjShtjliyLJmROSmzzjSYHjp8d14yESH3XK7A
chWP3qdySDVN/Lm3FYo4l5Z3zaf91kvrWTtg7nL4q9Qb2X0848FsKiJd//tziNDL
kYUCAwEAAaMgMB4wDgYDVR0PAQH/BAQDAgWgMAwGA1UdEwEB/wQCMAAwDQYJKoZI
hvcNAQELBQADggIBABeSzp7kfPoiie8Kq8MRkswwEV4hGUYlS+OQS0SCIrtCsbld
soY1M1SD0J3BEK9rNx4VMLV7BjPyluf4fuVdvqpX3feOFZfCtLgb/1jTtJ/YN9vs
28AezwUWRbNStp2+lw7y591iJcCfMU29lA65FmumJ44H95gFv1yku70ZsFhWY/c7
DIxBkvUEYCT8OQ3r63VZlI3jiiIU74VKVYdjTtpfTRYXxmKqwM9lGlLotr/ECUiG
/xuJOWQLx4ktOHbSgsKT1Zz1jpq6D6l2YyUrzCllhJ4qkp4FW6NpxibaCjhH1FWO
5MtS6on7YfdGSon1e9dGeiuLq0odJncK+lZcxNFLBmMGjwWY6Os7BqLFpG+x/97E
6eleiCbeAeqHtoMJMz7zeioSdJWxD7PA4RlwqXKjKHBkTQuhcpF+51B93Fo/tBfm
jrOJXNvfYbCm0vofKemy0MSzmB+1ZiY3dfuO6ZvrRHvg7PrmI2kA4IUY4FJX4O9+
rdI47qXcorftBmWUM5wpGkfBfOTJsDtIGWopTtWE28pSn1XVjOR0GZ1d2UJzbRcK
X9cs6msjO+y7zK4hR7cNwhwTZXa1w5BSA//vWgTvxzRDutDMIYHEBOMKB5i580hF
P5FUh5oNnrBCuyewtfmOjQgRhXT9YF2711luXqTFsqzLcLg9UbDazjs52cjq
-----END CERTIFICATE-----`;

function injectStyles() {
    GM.addStyle(`
        .jetbra-button {
            background-color: #04AA6D; border: none; color: white; padding: 8px 24px;
            text-align: center; text-decoration: none; display: inline-block;
            border-radius: 16px; box-shadow: 0 8px 16px 0 rgba(0,0,0,0.2), 0 6px 20px 0 rgba(0,0,0,0.19);
            transition-duration: 0.4s;
            margin-left:5px;
        }
        .jetbra-button:hover { background-color: #057e47; color: white; }
    `);
}


async function findElementWithRetry(cssSelector) {
    const maxAttempts = 50;
    for (let attempts = 0; attempts < maxAttempts; attempts++) {
        const element = document.querySelector(cssSelector);
        if (element) {
            return element;
        }
        await new Promise(resolve => setTimeout(resolve, 100));
    }
    throw new Error(`Element with selector '${cssSelector}' not found after ${maxAttempts} attempts.`);
}

function pem2base64(pem) {
    return pem.split('\n').reduce((base64, line) => line.includes("--") ? base64 : base64 + line, '');
}

function arrayBufferToBase64(buffer) {
    return btoa([...new Uint8Array(buffer)].map(b => String.fromCharCode(b)).join(''));
}

function base64ToArrayBuffer(base64) {
    return Uint8Array.from(atob(base64), c => c.charCodeAt(0)).buffer;
}

function genLicenseId() {
    const CHARSET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
    return Array.from({length: 10}, () => {
        let idx = Math.floor(Math.random() * CHARSET.length);
        return CHARSET[idx];
    }).join('');
}

function buildLicensePartJson(productCode, licenseId,twoYears) {
    let formattedDateTwoYears = '2099-08-01';
    if (twoYears) {
        let futureDateTwoYears = new Date();
        futureDateTwoYears.setFullYear(futureDateTwoYears.getFullYear() + 2);
        formattedDateTwoYears = futureDateTwoYears.toISOString().split('T')[0];
    }

    let fallbackDate = formattedDateTwoYears;
    let paidUpTo = formattedDateTwoYears;

    const products = [productCode].map(code => ({
            code: code,
            fallbackDate: fallbackDate,
            paidUpTo: paidUpTo
    }));

    return JSON.stringify({
        "licenseId": licenseId,
        "licenseeName": "cHAuHaN",
        "assigneeName": "cHAuHaN",
        "assigneeEmail": "",
        "licenseRestriction": "",
        "checkConcurrentUse": false,
        "products": products,
        "metadata": "0120230102PPAA013009",
        "hash": "41472961/0:1563609451",
        "gracePeriodDays": 7,
        "autoProlongated": true,
        "isAutoProlongated": true
    });
}

async function addButton() {
    injectStyles();

    let url = window.location.href
    if (!url.startsWith('https://plugins.jetbrains.com/plugin/')) {
        return;
    }

    let pluginId = url.split('/')[4].split('-')[0]
    console.log('pluginId: ' + pluginId);

    let pluginDetail = await fetch('https://plugins.jetbrains.com/api/plugins/' + pluginId).then(r => r.json());

    const parentElement = await findElementWithRetry('.plugin-header__controls-panel > div:first-child');

    if (parentElement.querySelector('.jetbra-button')) {
        return;
    }
    let newElement = document.createElement('div');
    newElement.classList.toggle('wt-col-inline');
    newElement.innerHTML =
        `<div class="customize-btn" style="display: flex; flex-direction: column;white-space: nowrap;">
            <div class="generate-plugin-code" style="display: flex;height: 40px;">
                <button class="jetbra-button" type="button" id="permanentLicense">Get permanent activation code</button>
                <button class="jetbra-button" type="button" id="twoYearsLicense">Get a 2-year activation code</button>
            </div>
        </div>`;

    parentElement.appendChild(newElement)

    newElement.addEventListener('click', async (e) => {
        let productCode = pluginDetail?.purchaseInfo?.productCode;
        if (productCode === undefined) {
            window.alert('This plugin is not a paid plugin');
            return;
        }
        let licenseId = genLicenseId()
        let licensePartJson = buildLicensePartJson(productCode, licenseId,e.target.id === "twoYearsLicense")

        let privateKey = await window.crypto.subtle.importKey("pkcs8", base64ToArrayBuffer(pem2base64(pemEncodedKey)), {
            name: "RSASSA-PKCS1-v1_5", hash: "SHA-1",
        }, true, ["sign"]);

        let licensePartBase64 = btoa(unescape(encodeURIComponent(licensePartJson)));
        let sigResultsBase64 = arrayBufferToBase64(await window.crypto.subtle.sign("RSASSA-PKCS1-v1_5", privateKey, new TextEncoder().encode(licensePartJson)));
        let cert_base64 = pem2base64(pemEncodedCrt);

        GM.setClipboard(`${licenseId}-${licensePartBase64}-${sigResultsBase64}-${cert_base64}`, 'text');
        window.alert('The activation code has been copied to your clipboard.');
    })
}

addButton();
if (window.onurlchange === null) {
    window.addEventListener('urlchange', (ignore) => {
        addButton();
    });
}