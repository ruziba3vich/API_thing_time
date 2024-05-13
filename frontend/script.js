function submitForm() {
    var form = document.getElementById("myForm");
    document.getElementById("submitButton").disabled = true;

    var formData = {
        day_of_week: form.elements["day_of_week"].value,
        day_of_month: parseInt(form.elements["day_of_month"].value),
        month: form.elements["month"].value,
        year: parseInt(form.elements["year"].value),
        hour: parseInt(form.elements["hour"].value),
        minute: parseInt(form.elements["minute"].value),
        second: parseInt(form.elements["second"].value)
    };

    fetch("http://localhost:7777/time/RFC3339", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(formData)
    })
    .then(response => response.json())
    .then(data => {
        document.getElementById("result").innerText = "Result: " + data.result;
    })
    .catch(error => {
        console.error("Error:", error);
        document.getElementById("errorMessage").innerText = "An error occurred while submitting the form.";
    });
}

var inputs = document.querySelectorAll("input");
inputs.forEach(function(input) {
    input.addEventListener("input", function() {
        var isValid = true;
        inputs.forEach(function(input) {
            if (!input.checkValidity()) {
                isValid = false;
            }
        });
        document.getElementById("submitButton").disabled = !isValid;
    });
});
