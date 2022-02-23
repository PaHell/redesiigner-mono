<script context="module" lang="ts">
	import Btn from '$lib/Button.svelte';
import { createEventDispatcher } from 'svelte';
	import { menuStore } from '../store';
	import Icon from "./Icon.svelte";
</script>

<script lang="ts">
	export let options : string[] = [];
	export let selected = 0;
	export let opened = false;
	export let bg : string | undefined;
	export let color : string | undefined;

	let ref : HTMLElement | undefined;

	const dispatch = createEventDispatcher();

	function toggleOpen() : void {
		opened = !opened;
		if (opened) {
			dispatch("open");
			menuStore.open(() => { opened = false });
		} else {
			dispatch("close");
			menuStore.open(() => { });
		}
	}
</script>

<template>
	<div class="app-select">
		<Btn active={opened} disableLoading
			bg={bg}
			color={color}
			on:click={toggleOpen}>
			<Icon name="home"/>
			<slot name="active" item={options[selected]}/>
			<Icon name="select"/>
		</Btn>
		<menu bind:this={ref}
			class="app-select-menu"
			class:open={opened}>
			<div>
				<div>
					{#each options as option, index}
						<Btn disableLoading
							color={selected === index ? 'accent' : 'default' }
							active={selected === index}
							bg="transparent"
							on:click={() => { selected = index; toggleOpen(); }}>
							<slot name="item"
								item={option}
								active={index == selected}/>
						</Btn>
					{/each}
				</div>
			</div>
		</menu>
	</div>
</template>

<style global lang="postcss">
	.app-select {
		@apply flex flex-col justify-start;
		min-width: 160px;

		& > .app-button {
			&.app-button-active {
				&.app-button-bg-transparent {
					@apply border-gray-300;
					&:after { content: none; }
				}
			}
			& > main {
				@apply w-full;
				& p {
					@apply flex-1 text-left;
				}
				& .app-icon {
					&:last-child {
						@apply indent-2 border-l;
					}
				}
			}
		}
		
	}

	.app-select-menu {
		@apply w-full h-0
		mt-[-5px]
		overflow-visible;
		margin-top: calc(-8px - 2px);
		
		& > div {
			@apply flex max-h-0 overflow-hidden;
			width: calc(100% + 2 * 8px);
			margin-left: -8px;
			transition: max-height .25s;
			will-change: max-height height;

			& > div {
				@apply relative overflow-hidden
				flex flex-col justify-center
				py-2 border border-gray-300
				bg-gray-50 shadow rounded-md;
				width: calc(100% - 2 * 8px);
				margin: 8px;

				& .app-button {
					@apply items-stretch flex-shrink-0 
					rounded-none transition-colors;
					@apply border-x-0 !important;

					& main {
						@apply justify-start;
					}
					
					&.app-button-active {
						&:before {
							content: "";
							@apply block absolute h-5 w-[2px] rounded-full;
						}
	
						&:after { @apply hidden; }
					}
				}
			}

		}
		&.open > div {
			@apply max-h-96;
		}
	}

	.dark .app-select {

		& > .app-button {
			&.app-button-active {
				&.app-button-bg-transparent {
					@apply border-gray-700;
				}
			}
			&:not(.app-button-bg-white) {
				& main .app-icon {
					&:last-child {
						@apply border-l-2;
						height: calc(100% + 1.25rem);
					}
				}
			}
		}
	}

	.dark .app-select-menu {
		margin-top: -6px;

		& > div > div {
			@apply bg-gray-800 shadow-none
			ring-2 ring-gray-900
			border-2 border-gray-700;
		}
	}

</style>
