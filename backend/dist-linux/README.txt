bai-jin-yan-api Linux package 
 
Usage: 
1. Upload the whole dist-linux directory to your Linux server 
2. Make sure manifest/config/config.yaml has the correct database host 
3. Run: chmod +x bai-jin-yan-api-linux-amd64 
4. Start: nohup ./bai-jin-yan-api-linux-amd64 > app.log 2>&1 & 
5. Check: tail -n 100 app.log 
