# API client for text-processing.com

This is a small wrapper / client for the natural language processing API's found at 
[text-processing.com](http://text-processing.com/).

The website provides 4 API's:

1. [Stemming and Lemmatization API](#stemming)
2. [Sentiment Analysis API](#analysis)
3. [Tagging and Chunk Extraction API](#tagging)
4. [Phrase Extraction and Named Entity Recognition API](#extraction)

The original API is documented at: http://text-processing.com/docs/index.html

## Stemming and Lemmatization API<a name="stemming"></a>

API documented at: http://text-processing.com/docs/phrases.html

**Status:** Not implemented yet

## Sentiment Analysis API<a name="analysis"></a>

API documented at: http://text-processing.com/docs/phrases.html

**Status:** Not implemented yet

## Tagging and Chunk Extraction API<a name="tagging"></a>

API documented at: http://text-processing.com/docs/phrases.html

**Status:** Not implemented yet

## Phrase Extraction and Named Entity Recognition API<a name="extraction"></a>

API documented at: http://text-processing.com/docs/phrases.html

### Example

This example comes from `phrase_test.go`:

```
result, err := ExtractPhrase("John hit the ball.")
```

Then `result` contains:

```
map[
    DATE:[ball] 
    PERSON:[John] 
    VP:[hit] 
    LOCATION:[John]
]
```

## License

The MIT License (MIT). Please see the `LICENSE` file for more information.