package prompt

import (
	"crypto/sha512"
	"encoding/hex"
	"strings"
)

func Createprompt(userQuery string) string {

	const blockInstruction string = "   - follow block structure very precisely and strictly"

	var sb strings.Builder

	sb.WriteString("# ANSWER the QUERY below" + "\n")
	sb.WriteString("\n")
	sb.WriteString(userQuery)
	sb.WriteString("To be able to answer the Query, you must follow the structure below including four blocks RESPONSE, REPHRASED, KEYWORDS and SUMMARY" + "\n")
	sb.WriteString("each block distinguished with specific starting and ending tags, the tags are explained in the following instructions" + "\n")
	sb.WriteString("No other extra terms or elements apart from defined tags should not be added to the answer." + "\n")
	sb.WriteString("\n")
	sb.WriteString("1. Answer to the query above should be set inside the following provided RESPONSE block." + "\n")
	sb.WriteString("   - The answer should be very elaborative" + "\n")
	sb.WriteString("   - pay attention if the user ask you to elaborate the answer in certain level such as age, school or university background" + "\n")
	sb.WriteString("   - You should match your elaboration with the user age, school and univeristy background if possible" + "\n")
	sb.WriteString("   - RESPONSE block starts with the tag " + "\n" + ResponseStartSha512 + "\n")
	sb.WriteString("   - RESPONSE block ends with the tag " + "\n" + ResponseEndSha512 + "\n")
	sb.WriteString(blockInstruction)
	sb.WriteString("\n" + ResponseStartSha512 + "\n")
	sb.WriteString("your answer comes here")
	sb.WriteString("\n" + ResponseEndSha512 + "\n")
	sb.WriteString("\n")

	sb.WriteString("2. Paraphrase the text in the RESPONSE block and set the paraphrase in the PARAPHRASED block" + "\n")
	sb.WriteString("   - make the paraphrase text as relevant as possible to the text in the RESPONSE block, and" + "\n")
	sb.WriteString("   - avoid jargon words" + "\n")
	sb.WriteString("   - start the summary with impretive verb form" + "\n")
	sb.WriteString("   - REPHRASED block starts with the tag " + "\n" + RephrasedStartSha512 + "\n")
	sb.WriteString("   - REPHRASED block ends with the tag " + "\n" + RephrasedEndSha512 + "\n")
	sb.WriteString(blockInstruction)
	sb.WriteString("\n" + RephrasedStartSha512 + "\n")
	sb.WriteString("your paraphrased text come here")
	sb.WriteString("\n" + RephrasedEndSha512 + "\n")
	sb.WriteString("\n")

	sb.WriteString("3. Collect the crucial terms or keywords selected carefully from the the RESPONSE block and REPHRASE block. the selected Keywords should be precise," + "\n")
	sb.WriteString("   - the keywords should make you enable to reconstruction of the RESPONSE block's gist or providing a clear understanding of RESPONSE block core topics." + "\n")
	sb.WriteString("   - elements can be common terms, symbols, etc." + "\n")
	sb.WriteString("   - the selected keywords should be preseted in a comma-separated string inside the following provided KEYWORDS block" + "\n")
	sb.WriteString("   - KEYWORDS block starts with the tag " + "\n" + KeywordsStartSha512 + "\n")
	sb.WriteString("   - KEYWORDS block ends with the tag " + "\n" + KeywordsEndSha512 + "\n")
	sb.WriteString(blockInstruction)
	sb.WriteString("\n" + KeywordsStartSha512 + "\n")
	sb.WriteString("your keywords presented here as a comma separated string")
	sb.WriteString("\n" + KeywordsEndSha512 + "\n")
	sb.WriteString("\n")

	sb.WriteString("4. Provide a concise summary of the text in RESPONSE block, and set it in the SUMMARY block" + "\n")
	sb.WriteString("   - with a maximum length of 1000 characters." + "\n")
	sb.WriteString("   - The summary should be as close as possible to RESPONSE and KEYWORDS blocks" + "\n")
	sb.WriteString("   - capturing the core gist of should RESPONSE and KEYWORDS blocks be highly focused." + "\n")
	sb.WriteString("   - SUMMARY block starts with the tag " + "\n" + SummaryStartSha512 + "\n")
	sb.WriteString("   - SUMMARY block ends with the tag " + "\n" + SummaryEndSha512 + "\n")
	sb.WriteString(blockInstruction)
	sb.WriteString("\n" + SummaryStartSha512 + "\n")
	sb.WriteString("your summary presented here, write it in a way that you can understand it yourself if I need to reminde you again, avoid jargon words." + "\n")
	sb.WriteString("This is summary is going to be used in future that you remember the gist of our conversation" + "\n")
	sb.WriteString("\n" + SummaryEndSha512 + "\n")
	sb.WriteString("\n")

	sb.WriteString("Example of blocks:")
	sb.WriteString("\n" + ResponseStartSha512 + "\n")
	sb.WriteString("this is your response")
	sb.WriteString("\n" + ResponseEndSha512 + "\n")
	sb.WriteString("\n")
	sb.WriteString("\n" + RephrasedStartSha512 + "\n")
	sb.WriteString("this is your rephrase")
	sb.WriteString("\n" + RephrasedEndSha512 + "\n")
	sb.WriteString("\n")
	sb.WriteString("\n" + KeywordsStartSha512 + "\n")
	sb.WriteString("this is your keywords")
	sb.WriteString("\n" + KeywordsEndSha512 + "\n")
	sb.WriteString("\n")
	sb.WriteString("\n" + SummaryStartSha512 + "\n")
	sb.WriteString("this is your summary")
	sb.WriteString("\n" + SummaryEndSha512 + "\n")
	sb.WriteString("\n")

	return sb.String()
}

func convertTagToSha512(tag string) string {

	tagBytes := []byte(tag)

	hasher := sha512.New()
	hasher.Write(tagBytes)
	hashBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashBytes)
}
