(async () => {
    const go = new Go(); // Provided by wasm_exec.js
    const wasm = await WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject);

    go.run(wasm.instance);

    function createAccount() {
        const name = document.getElementById("account_name").value;
        const number = document.getElementById("account_number").value;
        const address = document.getElementById("address").value;
        const amount = parseFloat(document.getElementById("amount").value);
        const iban = document.getElementById("iban").value;
        const type = document.getElementById("type").value;
        window.createAccount(name, number, address, amount, iban, type);
    }

    function retrieveAccounts() {
        const page = parseInt(document.getElementById("page").value);
        const size = parseInt(document.getElementById("size").value);
        window.retrieveAccounts(page, size);
    }

    function healthCheck() {
        window.healthCheck();
    }

    document.querySelector("#create-account-btn").onclick = createAccount;
    document.querySelector("#retrieve-accounts-btn").onclick = retrieveAccounts;
    document.querySelector("#health-check-btn").onclick = healthCheck;
})();
