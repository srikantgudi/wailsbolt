<script>
  import logo from './assets/images/logo-universal.png'
  import {Greet, CreateBucket, GetValue, AddKeyValue} from '../wailsjs/go/main/App.js'

  let resultText = "Please enter your name below 👇"
  let buket = 'Mybucket'
  let keyvalue = ""
  let newvalue = ""
  let key = ""
  let name

  function greet() {
    Greet(name).then(result => resultText = result)
  }

  function getkeyvalue() {
    GetValue(buket, key).then(result => keyvalue = result)
  }

  function addnewvalue() {
    AddKeyValue(buket, key, newvalue)
  }

  function createbuket() {
    if (!CreateBucket(buket)) {
        alert('Error creating bucket')
    } else {
        alert('Bucket opened / created !!')
    }
  }
  
</script>

<main>
  <img alt="Wails logo" id="logo" src="{logo}">
  <!-- <div class="result" id="result">{resultText}</div>
  <div class="input-box" id="input">
    <input autocomplete="off" bind:value={name} class="input" id="name" type="text"/>
    <button class="btn" on:click={greet}>Greet</button>
  </div> -->
  <div>
    <h3>Set Bucket</h3>
    <input class="elem" type="text" bind:value={buket} />
    <button on:click={createbuket} class="btn">Create if not existing</button>
  </div>
  <div>
    <h3>GET VALUE for key:</h3>
    Key: <input class="elem" type="text" placeholder="key.." bind:value={key} /> &laquo; enter a key to get value
    <button class="btn" disabled={!buket} on:click={getkeyvalue}>Get Value</button>
    <div class="result">Value: {keyvalue}</div>
  </div>
  <div>
    <h3>ADD KEY,VALUE:</h3>
    Key: <input class="elem" type="text" placeholder="key.." bind:value={key} />
    Value: <input class="elem" type="text" placeholder="new value.." bind:value={newvalue} />
    <button class="btn"  on:click={addnewvalue}>Add Value</button>
  </div>
</main>

<style>
    #logo {
        display: block;
        width: 6rem;
        height: 6rem;
        margin: auto;
        padding: 10% 0 0;
        background-position: center;
        background-repeat: no-repeat;
        background-size: 100% 100%;
        background-origin: content-box;
    }

    .result {
        height: 20px;
        line-height: 20px;
        margin: 1.5rem auto;
    }

    .elem {
        padding: 0.25rem 0.5rem;
    }
    .btn {
        padding: 0.25rem 0.5rem;
        border: none;
        box-shadow: 0 0 4px #999;
        background-color: #eee;
        cursor: pointer;
    }
    .btn:hover {
        background-color: lightcyan;
    }
</style>
