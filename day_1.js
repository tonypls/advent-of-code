const fs = require("fs");
const { get } = require("http");

function getFile(filePath) {
  try {
    const data = fs.readFileSync(filePath, "utf8");
    return data;
  } catch (err) {
    console.error(err);
  }
}

function joinFirstAndLast(numbers) {
  return Number(numbers[0].toString() + numbers[numbers.length - 1].toString());
}

function getFirstAndLastNumbers(str) {
  if (!str) return [0];
  const matches = str.match(/\d/g);
  if (matches === null) return [0]; // Return [0] or some other default value if no numbers are found

  return matches.map(Number);
}

function getFirstAndLastNumbersAndWords(str) {
  const numbers = {
    1: "one",
    2: "two",
    3: "three",
    4: "four",
    5: "five",
    6: "six",
    7: "seven",
    8: "eight",
    9: "nine",
  };

  let firstNumber = 0;
  let lastNumber = 0;
  let firstIndex = str.length;
  let lastIndex = -1;

  for (const [digit, word] of Object.entries(numbers)) {
    let parts = str.split(word);
    if (parts.length > 1 && parts[0].length < firstIndex) {
      firstNumber = Number(digit);
      firstIndex = parts[0].length;
    }
    if (
      parts.length > 1 &&
      parts[parts.length - 1].length <= str.length - lastIndex
    ) {
      lastNumber = Number(digit);
      lastIndex = str.length - parts[parts.length - 1].length;
    }

    parts = str.split(digit);
    if (parts.length > 1 && parts[0].length < firstIndex) {
      firstNumber = Number(digit);
      firstIndex = parts[0].length;
    }
    if (
      parts.length > 1 &&
      parts[parts.length - 1].length <= str.length - lastIndex
    ) {
      lastNumber = Number(digit);
      lastIndex = str.length - parts[parts.length - 1].length;
    }
  }

  return [firstNumber, lastNumber];
}

try {
  const testData1 = getFile("./test_input_1.text").split("\n");
  const testData2 = getFile("./test_input_2.text").split("\n");
  const data = getFile("./input.text").split("\n");

  const test1 = testData1
    .map(getFirstAndLastNumbers)
    .map(joinFirstAndLast)
    .reduce((a, b) => a + b);

  const result1 = data
    .map(getFirstAndLastNumbers)
    .map(joinFirstAndLast)
    .reduce((a, b) => a + b);

  const test2 = testData2
    .map(getFirstAndLastNumbersAndWords)
    .map(joinFirstAndLast)
    .reduce((a, b) => a + b);
  const result2 = data
    .map(getFirstAndLastNumbersAndWords)
    .map(joinFirstAndLast)
    .reduce((a, b) => a + b);

  console.log("Results:", test1, result1, test2, result2);
} catch (err) {
  console.error(err);
}
