var numberOfDrumbuttons = document.querySelectorAll(".drum").length;

for (var i = 0; i <numberOfDrumbuttons; i++) {
    document.querySelectorAll(".drum")[i].addEventListener("click", function() {
        
        var buttonInnerHTML = this.innerHTML; 

        makeSound(buttonInnerHTML);

        buttonAnimation(buttonInnerHTML);

    });

}

document.addEventListener("keypress", function(event) {

    makeSound(event.key);
    buttonAnimation(event.key);
});

function makeSound(key) {

    switch (key) {
        case "w":
            var audio = new Audio("Audio/Won Da Mo (feat. Crayon, Bayanni, Magixx  LADIPOE)--671d135c187a90ab76b99eaf6e03f0ee.mp3");
            audio.play();
            break;

            case "a":
            var audio = new Audio("Audio/Asake_ft_Olamide_-_Amapiano.mp3");
            audio.play();
            break;

            case "s":
                var audio = new Audio("Audio/Seyi-Vibez-Dejavu-(TrendyBeatz.com).mp3");
                audio.play();
                break;

            case "d":
                var audio = new Audio("Audio/Coming for You (feat. A1 x J1)--257dde469d4bfd8d6beacd02df4b0f68.mp3");
                audio.play();
                break;

            case "j":
                var audio = new Audio("Audio/Juice Wrld - Can t Die.mp3");
                audio.play();
                break;

            case "k":
                var audio = new Audio("Audio/Kizz-Daniel-Shu-Peru-(TrendyBeatz.com).mp3");
                audio.play();
                break;

            case "l":
                var audio = new Audio("Audio/Luis Fonsi, Daddy Yankee - Despacito (Lyrics) ft. Justin Bieber-1-1-1.mp3");
                audio.play();
                break;

                default: console.log(buttonInnerHTML);
        
    }
}

function buttonAnimation(currentKey) {

   var activeButton = document.querySelector("." + currentKey);

   activeButton.classList.add("pressed");

   setTimeout(function(){
    activeButton.classList.remove("pressed");
   }, 100);
}