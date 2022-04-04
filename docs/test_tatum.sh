export API_TOKEN=1ab39066-9304-4f7a-adca-dd4ee10d165b_100
export URL_PATH=http://me
# export URL_PATH=https://api-eu1.tatum.io

for INDEX in {1..100}
do
    curl -XGET -s \
        --header "x-api-key: $API_TOKEN" \
        -o /dev/null -w "%{http_code} | $INDEX | BITCOIN INFO \n" \
        ${URL_PATH}/v3/bitcoin/info &
    curl -XGET -s --header "x-api-key: $API_TOKEN" \
        -o /dev/null -w "%{http_code} | $INDEX | TRX a \n" \
        ${URL_PATH}/v3/bitcoin/transaction/ebf2cc0448c9bf25593be90b43240289a035d1f299614b8cfae60cb8e4debe59 &
    curl -XGET -s --header "x-api-key: $API_TOKEN" \
        -o /dev/null -w "%{http_code} | $INDEX | TRX b \n" \
        ${URL_PATH}/v3/bitcoin/transaction/ebf2cc0448c9bf25593be90b43240289a035d1f299614b8cfae60cb8e4debe59 &
    curl -XGET -s --header "x-api-key: $API_TOKEN" \
        -o /dev/null -w "%{http_code} | $INDEX | TRX c \n" \
        ${URL_PATH}/v3/bitcoin/transaction/ebf2cc0448c9bf25593be90b43240289a035d1f299614b8cfae60cb8e4debe59 &
done
