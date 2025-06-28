<script lang="ts">
	import { goto } from '$app/navigation';
	import Button from '../ui/button/button.svelte';
	import { Dialog } from '../ui/dialog';
	import DialogContent from '../ui/dialog/dialog-content.svelte';
	import DialogDescription from '../ui/dialog/dialog-description.svelte';
	import DialogHeader from '../ui/dialog/dialog-header.svelte';
	import DialogTitle from '../ui/dialog/dialog-title.svelte';
	import Input from '../ui/input/input.svelte';
	import Label from '../ui/label/label.svelte';

	type prop = {
		open?: boolean;
		email?: string;
		password?: string;
	};

	let { open = $bindable(false), email, password }: prop = $props();
</script>

<Dialog bind:open>
	<DialogContent class="sm:max-w-md">
		<DialogHeader class="pb-4 text-center">
			<DialogTitle class="text-2xl font-bold text-gray-900">Hey Your Back</DialogTitle>
			<DialogDescription class="text-gray-600">Lets sign in to your account</DialogDescription>
		</DialogHeader>

		<form action="" method="post" class="space-y-4">
			<div class="space-y-2">
				<Label for="email" class="text-sm font-medium text-gray-700">Email Address</Label>
				<Input
					id="email"
					type="email"
					placeholder="Enter your email"
					class="h-11 border-gray-300 focus:border-gray-500 focus:ring-gray-500"
					bind:value={email}
				/>
			</div>
			<div class="space-y-2">
				<Label for="password" class="text-sm font-medium text-gray-700">Password</Label>
				<Input
					id="password"
					type="password"
					placeholder="Enter your password"
					class="h-11 border-gray-300 focus:border-gray-500 focus:ring-gray-500"
					bind:value={password}
				/>
			</div>
			<div class="pt-2">
				<Button class="h-11 w-full bg-gray-600 font-medium hover:bg-gray-700" type="submit"
					>Sign In</Button
				>
			</div>
			<div class="pt-2 text-center">
				<p class="text-sm text-gray-600">
					New Here?
					<button
						type="button"
						class="cursor-pointer font-medium text-gray-600 hover:text-gray-700"
						onclick={() => {
							goto('/register');
						}}
					>
						Join Us
					</button>
				</p>
			</div>
		</form>
	</DialogContent>
</Dialog>
