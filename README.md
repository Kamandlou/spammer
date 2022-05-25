# Spammer
Simple spammer written with golang

## How to use?
1. Turn on allow less secure apps in sender gmail
2. Clone this repository
3. Run ``go build main.go``
## Flags 
``from`` flag gets sender gmail.

``password`` flag gets sender gmail password.

``to`` flag gets receiver email.

``msg`` flag gets email message.

```bash
./main --from=sender@gmail.com --password=password --to=receiver@gmail.com --msg="email message"
```