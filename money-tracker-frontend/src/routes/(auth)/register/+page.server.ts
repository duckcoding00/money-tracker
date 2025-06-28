import { fail, type Actions } from '@sveltejs/kit';
import type { PageServerLoad } from '../register/$types';

export const load = (async () => {
	return {};
}) satisfies PageServerLoad;

export const actions: Actions = {
	register: async ({ request, fetch }) => {
		const form = await request.formData();

		const username = form.get('username') as string;
		const email = form.get('email') as string;
		const password = form.get('password') as string;
		const confirm_password = form.get('confirm_password') as string;

		if (confirm_password !== password) {
			return fail(400, { message: 'Password must same with password' });
		}

		const body = JSON.stringify({
			username,
			email,
			password
		});
		try {
			const response = await fetch('http://127.0.0.1:8080/api/v1/user', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json',
					Accept: 'application/json'
				},
				body: body
			});

			const result = await response.json();

			console.log(result);

			if (!response.ok) {
				return fail(response.status, result);
			}

			return {
				result
			};
		} catch (error) {
			return fail(500, { message: 'Server error occurred' });
		}
	}
};
