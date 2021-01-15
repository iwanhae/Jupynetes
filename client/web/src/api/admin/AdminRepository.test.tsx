import { AdminRepository } from './AdminRepository';

test('GET /admin/getUser', () => {
  let repository = new AdminRepository();
  let result = repository.getUser();
  expect(result).not.toBeNull();
});

test('POST /admin/template', () => {
  let repository = new AdminRepository();
  let result = repository.postAdminTemplate();
  expect(result).not.toBeNull();
});

test('POST /admin/quota', () => {
  let repository = new AdminRepository();
  let result = repository.postAdminQuota();
  expect(result).not.toBeNull();
});

test('POST /admin/quota/{userid}', () => {
  let repository = new AdminRepository();
  let result = repository.postUserQuota();
  expect(result).not.toBeNull();
});