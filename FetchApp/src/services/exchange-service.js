const https = require('https');
const cache = new Map();

function convertToUSD(priceInIDR, callback) {
    const cacheKey = `IDR_USD_${priceInIDR}`;
    const cachedValue = cache.get(cacheKey);
  
    if (cachedValue) {
        callback(cachedValue);
    } else {
        const url = `https://v6.exchangerate-api.com/v6/ed2ec4e8a73820df3b7dc7b5/latest/USD`;
    
        https.get(url, function(response) {
            let data = '';
    
            response.on('data', function(chunk) {
            data += chunk;
            });
    
            response.on('end', function() {
            try {
                const jsonData = JSON.parse(data);
                const conversionRate = jsonData["conversion_rates"]["IDR"];
                const priceInUSD = priceInIDR / conversionRate;
    
                const roundedValue = priceInUSD.toFixed(2);
                cache.set(cacheKey, roundedValue);
    
                callback(roundedValue);
            } catch (error) {
                console.error(error);
                callback(null);
            }
            });
        }).on('error', function(error) {
            console.error(error);
            callback(null);
        });
    }
}

module.exports = {
    convertToUSD
};
