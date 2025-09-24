const baseUrl = "http://localhost:8080/"

async function registerNewLongUrl(event) {
    event.preventDefault()
    let inputField = event.target.elements.longUrlInput
    let longUrl = inputField.value

    if (longUrl === ""){
        return;
    }
    
    inputField.value = ""

    let shortUrl;

    try {
        const res = await fetch(baseUrl + "register?url=" + longUrl);
        if(!res.ok){
            throw new Error()
        }
        shortUrl = await res.text()
    }catch (error){
        console.error("An error occured: ", error)
        return;
    }

    var shortUrlDisplay = document.getElementById("shortUrlDisplay")
    if (shortUrlDisplay === undefined){
        console.error("Could not access the shortUrlDisplay")
    }
    shortUrlDisplay.value = baseUrl + shortUrl
}
