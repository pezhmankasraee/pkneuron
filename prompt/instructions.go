package prompt

import (
	"crypto/sha512"
	"encoding/hex"
	"strings"
)

func Createprompt(userQuery string) string {

	const blockInstruction string = "   - follow block structure very precisely and strictly"

	var sb strings.Builder

	sb.WriteString("# The Goal of this request is to ANSWER the QUERY below" + "\n")
	sb.WriteString("\n")
	sb.WriteString(userQuery)
	sb.WriteString("To answer the Query, follow the structure below including four blocks RESPONSE, REPHRASED, KEYWORDS and SUMMARY" + "\n")
	sb.WriteString("each block distinguished with specific starting and ending tags, the tags are explained in the following instructions" + "\n")
	sb.WriteString("No other extra terms or elements apart from defined tags should not be added to the answer." + "\n")
	sb.WriteString("\n")
	sb.WriteString("1. RESPONSE block:" + "\n")
	sb.WriteString("   1.1. Answer to the query should be set inside the RESPONSE block." + "\n")
	sb.WriteString("   1.2. The answer should be extensively elaborative" + "\n")
	sb.WriteString("   1.3. Pay attention if the user ask you to elaborate the answer in certain level such as age, school or university background" + "\n")
	sb.WriteString("        Therefore, The elaboration should be matched with the user age, school and univeristy background if mentioned by the user" + "\n")
	sb.WriteString("   1.4. No limits in case of number of words" + "\n")
	sb.WriteString("   1.5. RESPONSE block starts with the tag: " + "\n" + ResponseStartSha512 + "\n")
	sb.WriteString("   1.6. RESPONSE block ends with the tag: " + "\n" + ResponseEndSha512 + "\n")
	sb.WriteString(blockInstruction)
	sb.WriteString("\n" + ResponseStartSha512 + "\n")
	sb.WriteString("your answer comes here")
	sb.WriteString("\n" + ResponseEndSha512 + "\n")
	sb.WriteString("\n")

	sb.WriteString("2. REPHRASED block:" + "\n")
	sb.WriteString("   2.1. Paraphrase the text in the Query, and set the paraphrased query in the PARAPHRASED block" + "\n")
	sb.WriteString("   2.2. The paraphrased query must be as relevant as possible to the text in the Query" + "\n")
	sb.WriteString("   2.3. Avoid jargon words" + "\n")
	sb.WriteString("   2.4. Do not extend the rephrased query more than 500 words" + "\n")
	sb.WriteString("   2.5. Pay attention to the Query's tone, if the query is a question, or if it is an instruction, etc, try to adapt the paraphrased query with the tone of the Query" + "\n")
	sb.WriteString("   2.6. REPHRASED block starts with the tag " + "\n" + RephrasedStartSha512 + "\n")
	sb.WriteString("   2.7. REPHRASED block ends with the tag " + "\n" + RephrasedEndSha512 + "\n")
	sb.WriteString(blockInstruction)
	sb.WriteString("\n" + RephrasedStartSha512 + "\n")
	sb.WriteString("your paraphrased text come here")
	sb.WriteString("\n" + RephrasedEndSha512 + "\n")
	sb.WriteString("\n")

	sb.WriteString("3. KEYWORDS block:" + "\n")
	sb.WriteString("   3.1. Collect the crucial terms or keywords selected carefully from the the RESPONSE block and REPHRASE block. the selected Keywords should be precise," + "\n")
	sb.WriteString("   3.2. The terms or keywords should represent the core concepts in the RESPONSE block perceily" + "\n")
	sb.WriteString("   3.3. To verify if you have selected the most relevant keywords and terms: " + "\n")
	sb.WriteString("        you should be able reconstruct the RESPONSE block's gist or provid a clear understanding of RESPONSE block core concepts based on the provided keywords if needed later" + "\n")
	sb.WriteString("   3.4. selected terms and keywords can be common, technical, professional terms, symbols, abbreviations etc." + "\n")
	sb.WriteString("   3.5. the selected keywords should be preseted in a comma-separated string inside the following provided KEYWORDS block" + "\n")
	sb.WriteString("   3.6. KEYWORDS block starts with the tag " + "\n" + KeywordsStartSha512 + "\n")
	sb.WriteString("   3.7. KEYWORDS block ends with the tag " + "\n" + KeywordsEndSha512 + "\n")
	sb.WriteString(blockInstruction)
	sb.WriteString("\n" + KeywordsStartSha512 + "\n")
	sb.WriteString("your keywords presented here as a comma separated string")
	sb.WriteString("\n" + KeywordsEndSha512 + "\n")
	sb.WriteString("\n")

	sb.WriteString("4. SUMMARY block: " + "\n")
	sb.WriteString("   4.1. Provide a concise summary of the text in the RESPONSE block, and set the provided summary text in the SUMMARY block" + "\n")
	sb.WriteString("   4.2. Do not extend the summary text more than 1000 words." + "\n")
	sb.WriteString("   4.3. The summary text should be as close as possible to the provided texts or keywords in the RESPONSE block and the KEYWORDS block" + "\n")
	sb.WriteString("   4.4. To verify how the summary text is accurate, you should be able to reconstruct the Query, and the text in RESPONSE block, and also guess the terms in the KEYWORDS block correctly." + "\n")
	sb.WriteString("   4.5. SUMMARY block starts with the tag " + "\n" + SummaryStartSha512 + "\n")
	sb.WriteString("   4.6. SUMMARY block ends with the tag " + "\n" + SummaryEndSha512 + "\n")
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
