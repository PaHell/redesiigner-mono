<script lang="ts" context="module">
      import Btn from "$lib/Button.svelte";
      import Slct from "$lib/Select.svelte";
	import { theme } from '$lib/../store';
</script>

<script lang="ts">
	const languages = ["Deutsch", "English", "Francais", "Espanol"];
	const themes = ["OS Pref", "Light", "Dark"];
	let selectedLang = 0;
	let selectedTheme = 0;
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
		<div class="flex flex-wrap justify-center space-2">
			<Slct options={languages}
				bind:selected={selectedLang}>
				<svelte:fragment slot="active" let:item>
					<p>{item}</p>
				</svelte:fragment>
				<svelte:fragment slot="item" let:item let:active>
					<p class:active>{item}</p>
				</svelte:fragment>
			</Slct>

                  <div class="dark">
				<Slct options={themes}
					bind:selected={selectedTheme}>
					<svelte:fragment slot="active" let:item>
						<p>{item}</p>
					</svelte:fragment>
					<svelte:fragment slot="item" let:item let:active>
						<p class:active>{item}</p>
					</svelte:fragment>
				</Slct>
		      </div>

		</div>
	</div>
</template>
