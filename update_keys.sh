#!/bin/bash

# Create directory if it doesn't exist
mkdir -p bin/.jetbra-free

# Run openssl command to convert the key
openssl pkcs8 -topk8 -inform PEM -outform PEM -nocrypt -in bin/.jetbra-free/jetbra.key -out jetbra_pkcs8.key

# Read the new key and certificate files
NEW_KEY=$(cat jetbra_pkcs8.key)
NEW_CERT=$(cat bin/.jetbra-free/jetbra.pem)

# Update the JavaScript file
sed -i "s|const pemEncodedKey = \`.*\`;|const pemEncodedKey = \`$NEW_KEY\`;|" "JetBrains Activation Code Generator.js"
sed -i "s|const pemEncodedCrt = \`.*\`;|const pemEncodedCrt = \`$NEW_CERT\`;|" "JetBrains Activation Code Generator.js"

# Clean up temporary file
rm jetbra_pkcs8.key

echo "Keys have been updated successfully!" 