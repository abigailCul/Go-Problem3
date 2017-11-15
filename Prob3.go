package main
 
 import ( 
"math/rand"
  "time"
  "fmt"
  "regexp"
  )
 
 var responses = []string{
  "I’m not sure what you’re trying to say. Could you explain it to me?",
  "How does that make you feel?",
  "Why do you say that?",
 }
 

 func ElizaResponse(input string) string {
	 if matched, _ := regexp.MatchString(`(?i).*\bfather\b.*`, input); matched{
	 return "Why dont you tell me about your Father?"
	 }
	 //(?i) case insensitive

   re := regexp.MustCompile("I am ([^.!?]*)[.!?]?")
	if re.MatchString(input){
  
   return re.ReplaceAllString(input, "How do you know you are $1?")

  }
 
  return responses[rand.Intn(len(responses))]
 
 }
 
 func main() {
  rand.Seed(time.Now().UTC().UnixNano())
 
  fmt.Println("People say I look like both my mother and father.")
  fmt.Println(ElizaResponse("People say I look like both my mother and father."))
 
  fmt.Println("\nFather was a teacher.")
  fmt.Println(ElizaResponse("Father was a teacher."))
 
  fmt.Println("\nI was my father’s favourite.")
  fmt.Println(ElizaResponse("I was my father’s favourite."))
 
  fmt.Println("\nI'm looking forward to the weekend.")
  fmt.Println(ElizaResponse("I'm looking forward to the weekend."))
 
  fmt.Println("\nMy grandfather was French!")
  fmt.Println(ElizaResponse("My grandfather was French!"))
 
  fmt.Println("\nI am happy.")
  fmt.Println(ElizaResponse("I am happy."))

 fmt.Println("\nI am not happy with your responses")
  fmt.Println(ElizaResponse("I am not happy with your responses"))

   fmt.Println("\nI am not sure that you understand the effect that your questions are having on me.")
  fmt.Println(ElizaResponse("I am not sure that you understand the effect that your questions are having on me."))

   fmt.Println("\nI am supposed to just take what you’re saying at face value?")
  fmt.Println(ElizaResponse("I am supposed to just take what you’re saying at face value?"))

 }
