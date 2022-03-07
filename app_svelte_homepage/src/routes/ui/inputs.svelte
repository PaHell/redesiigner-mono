<script lang="ts" context="module">
      import Btn from "$lib/Button.svelte";
      import TextInput from "$lib/TextInput.svelte";
      import Alert from "$lib/Alert.svelte";
	import { theme } from '$lib/../store';
</script>

<script lang="ts">
      let value = "Value";
</script>

<template>
	<div class="w-full p-4 mt-8 rounded bg-gray-50 dark:bg-gray-800 border-gray-300 dark:border-gray-700">
		<div class="app-bar-button justify-center">
			<p>Theme</p>
			<Btn icon="device" text="OS Pref"
				active={$theme === 'os'}
				on:click={() => theme.set('os')} />
			<Btn icon="device" text="Light"
				active={$theme === 'light'}
				on:click={() => theme.set('light')} />
			<Btn icon="device" text="Dark"
				active={$theme === 'dark'}
				on:click={() => theme.set('dark')} />
		</div>
		<div class="flex flex-col justify-center space-y-2">
                  <div class="flex justify-center space-x-2">
                        <TextInput
                              bind:value={value}
                              placeholder="Placeholder"/>
                        <TextInput
                              bind:value={value}
                              placeholder="Placeholder"
                              icon="home"/>
                        <TextInput
                              bind:value={value}
                              placeholder="Placeholder"
                              icon="home"
                              label="Email"
                              rules={["email"]}/>
                  </div>
                  <div class="flex justify-center space-x-2">
                        <TextInput
                              bind:value={value}
                              placeholder="Placeholder"
                              icon="home">
                              <Btn icon="cross" bg="transparent" />
                        </TextInput>
                        <TextInput
                              bind:value={value}
                              placeholder="Placeholder"
                              icon="home">
                              <Btn icon="cross" bg="accent" />
                        </TextInput>
                  </div>
		</div>
            <div class="flex flex-col justify-center space-y-2">
                  <Alert text="Error" message="An error occured processing your request." icon="error" color="default"/>
                  <Alert text="Error" message="An error occured processing your request." icon="error" color="accent"/>
                  <Alert text="Error" message="An error occured processing your request." icon="error" color="success"/>
                  <Alert text="Error" message="An error occured processing your request." icon="error" color="warning"/>
                  <Alert text="Error" message="An error occured processing your request." icon="error" color="danger"/>
                  <Alert text="Error" message="An error occured processing your request." icon="error" color="info"/>
            </div>
	</div>
</template>
