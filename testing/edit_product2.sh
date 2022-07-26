curl \
    -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTg4MTk3NTEsInJvbGUiOiJ1c2VyXG4iLCJ1c2VybmFtZSI6ImJyaWFuIn0.UlIPeryVfeZrwoveE6MBuMSs3f6WDrkjBfnurJJwDLI" \
    -d '{"test_message":"HELLO WORLD"}' \
    -X POST \
    http://127.0.0.1:8081/v1/api/product/edit_product \

