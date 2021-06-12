# Text Analysis API

This api is fully based on text analysis. It provides endpoints for 

1. Sentimental Analysis
2. A Word Dictionary
3. Translation 
4. Language Detection

# Endpoints & Query Parameters (WIP)

1. Sentimental Analysis

```
https://api.rhydderchc.rocks/api/v1/sentiments?sentence=TEXT
```
2. Dictionary

> The default language code is en_US

```
https://api.rhydderchc.rocks/api/v1/dictionary?word=WORD&lc=LANGUAGE_CODE

```
3. Translation

```
https://api.rhydderchc.rocks/api/v1/translate?text=TEXT&from=FROM&to=TO
```
4. Language Detection

```
https://api.rhydderchc.rocks/api/v1/detect_language?text=TEXT
```


