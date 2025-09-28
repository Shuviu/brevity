const baseUrl = "https://brevity.shuviu.de/api/"

async function registerNewLongUrl(event) {
    event.preventDefault()
    let inputField = event.target.elements.longUrlInput
    let longUrl = inputField.value

    if (longUrl === ""){
        alert("Please insert the URL to be shortened")
        return;
    }
    
    inputField.value = ""

    let shortUrl;

    try {
        const res = await fetch(baseUrl + "register?url=" + longUrl);
        if(!res.ok){
            alert("Oops.. something went wrong with registering your url.\nPlease contact the administrator")
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
