// Function to convert decimal number to Roman numeral
const inputField = document.getElementById('number');
const outputField = document.getElementById('output');
const convertButton = document.getElementById('convert-btn');

function convertToRoman() {
  let arabic = inputField.value;
  let roman = '';

  if (arabic === '') {
    outputField.innerHTML = 'Please enter a valid number.';
  } else if (arabic < 0) {
    outputField.innerHTML = 'Please enter a number greater than or equal to 1.';
  } else if (arabic > 3999) {
    outputField.innerHTML = 'Please enter a number less than or equal to 3999.';
  } else {
    // Rest of the conversion logic
    const arabicArray = [1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1];
    const romanArray = [
      'M',
      'CM',
      'D',
      'CD',
      'C',
      'XC',
      'L',
      'XL',
      'X',
      'IX',
      'V',
      'IV',
      'I',
    ];

    for (let i = 0; i < arabicArray.length; i++) {
      while (arabicArray[i] <= arabic) {
        roman += romanArray[i];
        arabic -= arabicArray[i];
      }
    }

    outputField.innerHTML = roman;
  }
}
 //Reload the page
function resetResult() {
  inputField.value = '';
  outputField.innerHTML = ''; 
  location.reload(); 
}