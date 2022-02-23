<script lang="ts" context="module">
	import Icon from '$lib/Icon.svelte';
	import LoadingSpinner from './LoadingSpinner.svelte';
	import { createEventDispatcher, onMount } from 'svelte';
</script>

<script lang="ts">
	export let text: string | undefined = undefined;
	export let icon: string | undefined = undefined;
	export let color: string | undefined = undefined;
	export let bg: string | undefined = undefined;
	export let disabled: boolean = false;
	export let loading: boolean = false;
	export let active: boolean = false;
	export let css: string | undefined = undefined;
	export let disableLoading = false;

	let ref: HTMLButtonElement | undefined;
	const dispatch = createEventDispatcher();

	onMount(() => {
		ref?.focus();
	});

	function onClick(e: Event): void {
		dispatch('click');
		if (disableLoading) return;
		loading = true;
		setTimeout(() => {
			loading = false;
		}, 500);
	}
</script>

<template>
	<button
		on:click={onClick}
		bind:this={ref}
		disabled={disabled||loading}
		style="outline: none;"
		class={`
			${css ?? ''}
			app-button
			${disabled ? 'app-button-disabled' : ''}
			${loading ? 'app-button-loading' : ''}
			${active ? 'app-button-active' : ''}
			app-button-bg-${bg ?? 'default'}
			app-button-color-${color ?? 'default'}
        	`}>
		{#if loading}
			<div class="animate__animated animate__bounceIn">
				<LoadingSpinner />
			</div>
		{/if}
		<main>
			{#if icon}
				<Icon name={icon} />
			{/if}
			<slot/>
			{#if text}
				<p>{text}</p>
			{/if}
		</main>
	</button>
</template>

<style global lang="postcss">
	.app-bar-button {
		@apply flex
		flex-col-reverse sm:flex-row
		sm:justify-end
		sm:p-3 p-4;

		&.app-bar-reverse {
			@apply sm:justify-start;
		}

		& > .app-button {
			&:not(:last-child) {
				@apply
				mt-2 sm:mt-0
				mr-0 sm:mr-2;
			}
		}

		& > p {
			@apply text-gray-500 px-2 self-center;
		}
	}

	.app-button {
		@apply flex flex-col items-center justify-center flex-shrink-0
        	h-10 sm:h-9 px-2 border rounded-md font-medium
        	transition-colors select-none;

		&.app-button-no-ring {
			@apply ring-0 ring-offset-0 !important;
		}

		& > div {
			@apply flex justify-center items-center h-0 absolute;
		}

		& .app-loading-spinner > div {
			@apply bg-accent-500;
		}

		& > div + main {
			@apply opacity-0;
		}

		& > main {
			@apply flex items-center justify-center
			transition-opacity border-inherit;
			& * {
				@apply border-inherit transition-colors;
			}
			& > .app-icon {
				@apply md:text-ic-sm;
				&:first-child:not(:last-child) { @apply ml-2; }
			}
			& > p {
				@apply px-2 mb-[.5px] font-medium
				text-base md:text-sm
				leading-ic-base md:leading-ic-sm;

				&:last-child:not(:first-child) {
					@apply pr-3;
				}
			}
		}
		
		/* DISABLED */

		&:disabled { @apply pointer-events-none; }

		&.app-button-disabled {
			&:not(.app-button-bg-white):not(.app-button-bg-transparent) {
				@apply shadow-none bg-gray-200 border-gray-300 !important;
				& p {  @apply text-gray-500 !important; }
				& .app-icon {  @apply text-gray-400 !important; }
			}
			&.app-button-bg-transparent {
				@apply border-transparent bg-transparent !important;
				& p {  @apply text-gray-500 !important; }
				& .app-icon {  @apply text-gray-400 !important; }
			}
			&.app-button-bg-white {
				box-shadow: none !important;
				& * { @apply text-inherit !important; }
				&.app-button-color-accent {
					@apply bg-accent-400 text-accent-700 !important;
				}
				&.app-button-color-success {
					@apply bg-success-400 text-success-700 !important;
				}
				&.app-button-color-warning {
					@apply bg-warning-400 text-warning-700 !important;
				}
				&.app-button-color-danger {
					@apply bg-danger-400 text-danger-700 !important;
				}
				&.app-button-color-info {
					@apply bg-info-400 text-info-700 !important;
				}
			}
		}

		/* TEXT COLORS */

		&.app-button-bg-white,
		&.app-button-bg-default,
		&.app-button-bg-transparent {
			& p:not(:first-child) {
				@apply text-gray-900;
			}
		}
		&.app-button-color-accent {
			@apply text-accent-600;
			&:hover,
			&:focus,
			&:active {
				@apply text-accent-700;
			}
		}
		&.app-button-color-success {
			@apply text-success-600;
			& .app-loading-spinner > div {
				@apply bg-success-500;
			}
			&:hover,
			&:focus,
			&:active {
				@apply text-success-700;
			}
		}
		&.app-button-color-warning {
			@apply text-warning-700;
			& .app-loading-spinner > div {
				@apply bg-warning-500;
			}
			&:hover,
			&:focus,
			&:active {
				@apply text-warning-800;
			}
		}
		&.app-button-color-danger {
			@apply text-danger-600;
			& .app-loading-spinner > div {
				@apply bg-danger-500;
			}
			&:hover,
			&:focus,
			&:active {
				@apply text-danger-700;
			}
		}
		&.app-button-color-info {
			@apply text-info-600;
			& .app-loading-spinner > div {
				@apply bg-info-500;
			}
			&:hover,
			&:focus,
			&:active {
				@apply text-info-700;
			}
		}

		/* BG COLORS */

		&.app-button-bg-accent {
			@apply bg-accent-600 border-accent-700;
			&:hover { @apply bg-accent-500 border-accent-600; }
			&:active,
			&:focus { @apply bg-accent-700 border-accent-800; }
		}
		&.app-button-bg-success {
			@apply bg-success-600 border-success-700;
			&:hover { @apply bg-success-500 border-success-600; }
			&:active,
			&:focus { @apply bg-success-700 border-success-800; }
		}
		&.app-button-bg-warning {
			@apply bg-warning-400 border-warning-500;
			&:hover { @apply bg-warning-300 border-warning-400; }
			&:focus,
			&:active { @apply bg-warning-500 border-warning-600; }
		}
		&.app-button-bg-danger {
			@apply bg-danger-600 border-danger-700;
			&:hover { @apply bg-danger-500 border-danger-600; }
			&:active,
			&:focus { @apply bg-danger-700 border-danger-800; }
		}
		&.app-button-bg-info {
			@apply bg-info-600 border-info-700;
			&:hover { @apply bg-info-500 border-info-600; }
			&:active,
			&:focus { @apply bg-info-700 border-info-800; }
		}

		/* BG VARIANTS */

		&.app-button-bg-white,
		&.app-button-bg-default {
			@apply bg-gray-50 shadow-sm;
			&:hover {
				@apply shadow bg-white;
			}
			&:active,
			&:focus {
				@apply shadow-inner bg-gray-100;
			}
		}

		&.app-button-bg-white {
			@apply shadow-md;

			& > main { @apply border-gray-300; }
			&:hover {
				@apply shadow-lg;
			}
			&:focus,
			&:active {
				@apply shadow-none;
			}

			&.app-button-color-accent {
				@apply border-accent-600;
			}
			&.app-button-color-success {
				@apply border-success-600;
			}
			&.app-button-color-warning {
				@apply border-warning-600;
			}
			&.app-button-color-danger {
				@apply border-danger-600;
			}
			&.app-button-color-info {
				@apply border-info-600;
			}
		}

		&.app-button-bg-default,
		&.app-button-bg-transparent {
			&.app-button-color-default {
				& p {
					@apply text-gray-700;
				}
				& .app-icon {
					@apply text-gray-500;
				}
				&:hover,
				&:focus,
				&:active {
					& p {
						@apply text-gray-900;
					}
					& .app-icon {
						@apply text-gray-700;
					}
				}
			}
		}

		&.app-button-bg-transparent {
			@apply bg-transparent border-transparent;

			& > div {
				@apply p-2
				bg-gray-50 border border-gray-300
				shadow-sm rounded;
			}
			&:hover { @apply bg-gray-100 border-gray-300; }
			&:focus,
			&:active { @apply bg-gray-200 border-gray-300; }

		}

		/* BG COLORS */

		/* white text on colored buttons only */
		&.app-button-bg-accent,
		&.app-button-bg-success,
		&.app-button-bg-danger,
		&.app-button-bg-info {
			@apply text-white;
			& .app-loading-spinner > div {
				@apply bg-white;
			}
		}
		/* except warning (orange) */
		&.app-button-bg-warning {
			@apply text-warning-900;
			& .app-loading-spinner > div {
				@apply bg-warning-800;
			}
		}

		/* ACTIVE */

		&.app-button-active {
			@apply relative;
			/* colored underline on transparent */
			&.app-button-bg-transparent {
				&:after {
					content: "";
					@apply block
					w-2/3 h-[2px] mt-[4px] mb-[-6px]
					bg-accent-500 rounded-full;
				}
				&.app-button-color-success:after { @apply bg-success-500; }
				&.app-button-color-warning:after { @apply bg-warning-400; }
				&.app-button-color-danger:after { @apply bg-danger-500; }
				&.app-button-color-info:after { @apply bg-info-500; }
			}
			/* ring colors for non transparent */
			&:not(.app-button-bg-transparent) {
				@apply ring-2 ring-offset-2
				ring-offset-gray-100 ring-accent-500;

				&.app-button-bg-success,
				&.app-button-color-success {
					@apply ring-success-500;
				}

				&.app-button-bg-warning,
				&.app-button-color-warning {
					@apply ring-warning-400;
				}

				&.app-button-bg-danger,
				&.app-button-color-danger {
					@apply ring-danger-500;
				}

				&.app-button-bg-info,
				&.app-button-color-info {
					@apply ring-info-500;
				}
			}
			&.app-button-bg-white {
				&.app-button-color-accent {
					@apply ring-accent-200 ring-offset-accent-500;
				}
				&.app-button-color-success {
					@apply ring-success-200 ring-offset-success-500;
				}
				&.app-button-color-warning {
					@apply ring-warning-300 ring-offset-warning-500;
				}
				&.app-button-color-danger {
					@apply ring-danger-200 ring-offset-danger-500;
				}
				&.app-button-color-info {
					@apply ring-info-200 ring-offset-info-500;
				}
			}
		}
	}

	.dark {
		& .app-bar-button {
			& > p {
				@apply text-gray-400;
			}
		}
		& .app-button {

			&:not(.app-button-bg-white) {
				@apply border-2;
				@apply shadow-none !important;
			}

			/* DISABLED */

			&.app-button-disabled {
				&:not(.app-button-bg-white):not(.app-button-bg-transparent) {
					@apply bg-gray-800 border-gray-900 !important;
					& p {  @apply text-gray-500 !important; }
					& .app-icon {  @apply text-gray-500 !important; }
				}
				&.app-button-bg-transparent {
					& p {  @apply text-gray-500 !important; }
					& .app-icon {  @apply text-gray-500 !important; }
				}
			}

			/* TEXT COLORS */

			&.app-button-bg-default,
			&.app-button-bg-transparent {
				& p:not(:first-child) {
					@apply text-white;
				}
			}

			&.app-button-color-accent {
				@apply text-accent-500;
				&:hover,
				&:focus,
				&:active {
					@apply text-accent-400;
				}
			}
			&.app-button-color-success {
				@apply text-success-500;
				&:hover,
				&:focus,
				&:active {
					@apply text-success-400;
				}
			}
			&.app-button-color-warning {
				@apply text-warning-400;
				&:hover,
				&:focus,
				&:active {
					@apply text-warning-500;
				}
			}
			&.app-button-color-danger {
				@apply text-danger-500;
				&:hover,
				&:focus,
				&:active {
					@apply text-danger-400;
				}
			}
			&.app-button-color-info {
				@apply text-info-500;
				&:hover,
				&:focus,
				&:active {
					@apply text-info-400;
				}
			}

			/* BG COLORS */

			&.app-button-bg-accent {
				@apply bg-accent-600 border-accent-600;
				&:hover { @apply bg-accent-600 border-accent-500; }
				&:active,
				&:focus { @apply bg-accent-500 border-accent-500; }
			}
			&.app-button-bg-success {
				@apply bg-success-600 border-success-600;
				&:hover { @apply bg-success-600 border-success-500; }
				&:active,
				&:focus { @apply bg-success-500 border-success-500; }
			}
			&.app-button-bg-warning {
				@apply bg-warning-500 border-warning-500;
				&:hover { @apply bg-warning-500 border-warning-400; }
				&:focus,
				&:active { @apply bg-warning-400 border-warning-400; }
			}
			&.app-button-bg-danger {
				@apply bg-danger-600 border-danger-600;
				&:hover { @apply bg-danger-600 border-danger-500; }
				&:active,
				&:focus { @apply bg-danger-500 border-danger-500; }
			}
			&.app-button-bg-info {
				@apply bg-info-600 border-info-600;
				&:hover { @apply bg-info-600 border-info-500; }
				&:active,
				&:focus { @apply bg-info-500 border-info-500; }
			}

			/* BG VARIANTS */

			&.app-button-bg-default {
				@apply bg-gray-700 border-gray-700;
				&:hover {
					@apply bg-gray-700 border-gray-600;
				}
				&:active,
				&:focus {
					@apply bg-gray-600 border-gray-600;
				}
			}

			&.app-button-bg-default,
			&.app-button-bg-transparent {
				&.app-button-color-default {
					& p {
						@apply text-gray-300;
					}
					& .app-icon {
						@apply text-gray-400;
					}
					&:hover,
					&:focus,
					&:active {
						& p {
							@apply text-gray-100;
						}
						& .app-icon {
							@apply text-gray-200;
						}
					}
				}
			}

			&.app-button-bg-transparent {
				& > div {
					@apply bg-gray-900 border-gray-900;
				}
				&:hover { @apply bg-gray-800 border-gray-700; }
				&:focus,
				&:active { @apply bg-gray-700 border-gray-700; }
			}

			/* BG COLORS */

			/* ACTIVE */

			&.app-button-active {
				/* colored underline on transparent */
				&.app-button-bg-transparent {
					&:after {
						@apply bg-accent-500;
					}
					&.app-button-color-success:after { @apply bg-success-500; }
					&.app-button-color-warning:after { @apply bg-warning-400; }
					&.app-button-color-danger:after { @apply bg-danger-500; }
					&.app-button-color-info:after { @apply bg-info-500; }
				}
				/* ring colors for non transparent */
				&.app-button-bg-default {
					@apply bg-gray-900
					ring-0 ring-offset-0 border-accent-500;

					&.app-button-color-accent {
						@apply border-accent-500;
					}
					&.app-button-color-success {
						@apply border-success-500;
					}
					&.app-button-color-warning {
						@apply border-warning-400 text-warning-400 !important;
					}
					&.app-button-color-danger {
						@apply border-danger-500;
					}
					&.app-button-color-info {
						@apply border-info-500;
					}
				}
				&.app-button-bg-accent,
				&.app-button-bg-success,
				&.app-button-bg-warning,
				&.app-button-bg-danger,
				&.app-button-bg-info {
					@apply ring-offset-0 ring-0 border-white;
				}
			}
		}
	}
</style>
