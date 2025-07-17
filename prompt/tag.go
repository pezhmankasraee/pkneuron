package prompt

const responseStart string = "{{<PKNEURON::RESPONSE"
const responseEnd string = "\nPKNEURON::RESPONSE>}}\n"
const rephrasedStart string = "\n{{<PKNEURON::REPHRASE\n"
const rephrasedEnd string = "\nPKNEURON::REPHRASE>}}\n"
const keywordsStart string = "\n{{<PKNEURON::KEYWORDS\n"
const keywordsEnd string = "\nPKNEURON::KEYWORDS>}}\n"
const summaryStart string = "\n{{<PKNEURON::SUMMARY\n"
const summaryEnd string = "\nPKNEURON::SUMMARY>}}\n"

var ResponseStartSha512 string = convertTagToSha512(responseStart)
var ResponseEndSha512 string = convertTagToSha512(responseEnd)
var RephrasedStartSha512 string = convertTagToSha512(rephrasedStart)
var RephrasedEndSha512 string = convertTagToSha512(rephrasedEnd)
var KeywordsStartSha512 string = convertTagToSha512(keywordsStart)
var KeywordsEndSha512 string = convertTagToSha512(keywordsEnd)
var SummaryStartSha512 string = convertTagToSha512(summaryStart)
var SummaryEndSha512 string = convertTagToSha512(summaryEnd)
