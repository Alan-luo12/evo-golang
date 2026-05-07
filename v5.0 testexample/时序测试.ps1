
$submit = Invoke-RestMethod `
  -Uri http://localhost:8080/Submit `
  -Method Post `
  -ContentType "application/json" `
  -Body '{"name":"ttl_case_1","delay_time":5000}'

$submit | ConvertTo-Json -Depth 5

$id = $submit.data.task_id
Write-Host "task_id = $id"

Write-Host "========== 立刻查询 =========="
Invoke-RestMethod "http://localhost:8080/Getstatus?id=$id" | ConvertTo-Json -Depth 5

Start-Sleep -Seconds 1
Write-Host "========== 1秒后查询 =========="
Invoke-RestMethod "http://localhost:8080/Getstatus?id=$id" | ConvertTo-Json -Depth 5

Start-Sleep -Seconds 310
Write-Host "========== TTL过后查询（5分钟后） =========="
try {
    Invoke-RestMethod "http://localhost:8080/Getstatus?id=$id" | ConvertTo-Json -Depth 5
} catch {
    Write-Host "请求失败："
    $_
}
