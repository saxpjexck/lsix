# Create directory if it doesn't exist
New-Item -ItemType Directory -Force -Path "bin\.jetbra-free"

# Run openssl command to convert the key
openssl pkcs8 -topk8 -inform PEM -outform PEM -nocrypt -in "bin\.jetbra-free\jetbra.key" -out "jetbra_pkcs8.key"

# Read the new key and certificate files
$NEW_KEY = Get-Content -Path "jetbra_pkcs8.key" -Raw
$NEW_CERT = Get-Content -Path "bin\.jetbra-free\jetbra.pem" -Raw

# Escape special characters for regex
$NEW_KEY = [regex]::Escape($NEW_KEY)
$NEW_CERT = [regex]::Escape($NEW_CERT)

# Update the JavaScript file
$jsContent = Get-Content -Path "JetBrains Activation Code Generator.js" -Raw
$jsContent = $jsContent -replace "const pemEncodedKey = `.*`;", "const pemEncodedKey = `$NEW_KEY`;"
$jsContent = $jsContent -replace "const pemEncodedCrt = `.*`;", "const pemEncodedCrt = `$NEW_CERT`;"
Set-Content -Path "JetBrains Activation Code Generator.js" -Value $jsContent

# Clean up temporary file
Remove-Item -Path "jetbra_pkcs8.key" -Force

Write-Host "Keys have been updated successfully!" 