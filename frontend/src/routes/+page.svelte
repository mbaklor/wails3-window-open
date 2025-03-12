<script>
    import { GreetService } from "../../bindings/changeme";
    import { Events } from "@wailsio/runtime";

    let name = "";
    let result = "Please enter your name below 👇";
    let time = "Listening for Time event...";

    let openCount = 0;

    const doGreet = () => {
        let localName = name;
        if (!localName) {
            localName = "anonymous";
        }
        GreetService.Greet(localName)
            .then((resultValue) => {
                result = resultValue;
            })
            .catch((err) => {
                console.log(err);
            });
    };

    const settings = async () => {
        openCount = openCount + 1;
        await GreetService.OpenSettings();
    };

    Events.On("time", (timeValue) => {
        time = timeValue.data;
    });
</script>

<div class="container">
    <div>
        <span data-wml-openURL="https://wails.io">
            <img src="/wails.png" class="logo" alt="Wails logo" />
        </span>
        <span data-wml-openURL="https://svelte.dev">
            <img src="/svelte.svg" class="logo svelte" alt="Svelte logo" />
        </span>
    </div>
    <h1>Wails + Svelte</h1>
    <div class="result">{result}</div>
    <div class="card">
        <div class="input-box">
            <input
                class="input"
                bind:value={name}
                type="text"
                autocomplete="off"
            />
            <button class="btn" on:click={doGreet}>Greet</button>
            <button class="btn" on:click={settings}>Settings</button>
        </div>
    </div>
    <div class="footer">
        <div><p>Click on the Wails logo to learn more</p></div>
        <div><p>{time}</p></div>
        <div><p>Settings page has been opened {openCount} times!</p></div>
    </div>
</div>

<style>
</style>
