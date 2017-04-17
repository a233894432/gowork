function greetDevelopers(list) {
    return list.map(x => (x.greeting = `Hi ${x.firstName}, what do you like the most about ${x.language}?`, x));
}
var na = greetDevelopers([{
    firstName: "diogo",
    language: 'golang'
}])
console.log(na)