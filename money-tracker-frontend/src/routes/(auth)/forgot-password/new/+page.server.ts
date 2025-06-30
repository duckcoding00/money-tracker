import { fail, isRedirect, redirect, type Actions } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load = (async () => {
	return {};
}) satisfies PageServerLoad;

export const actions: Actions = {
	update: async ({ fetch, request }) => {
		const form = await request.formData();
		const username = form.get('username') as string;
		const token = form.get('token') as string;
		const password = form.get('password') as string;
		const confirm_password = form.get('confirm_password') as string;

		if (confirm_password !== password) {
			return fail(400, { message: 'Password must same with confirm password' });
		}

		const body = JSON.stringify({
			password
		});

		try {
			const response = await fetch(
				`http://127.0.0.1:8080/api/v1/user/reset-password?token=${token}&username=${username}`,
				{
					method: 'PATCH',
					headers: {
						'Content-Type': 'application/json',
						Accept: 'application/json'
					},
					body: body
				}
			);

			const result = await response.json();

			if (!response.ok) {
				return fail(response.status, result);
			}

			throw redirect(303, '/login');
		} catch (error) {
			if (isRedirect(error)) throw error;
			return fail(500, { message: 'Server error occurred' });
		}
	}
};
