from playwright.sync_api import sync_playwright, expect

def test_logged_out_flow(page):
    # 1. Navigate to the dashboard and verify it's protected
    page.goto("http://localhost:5173/dashboard")

    # Expect to see the "Access Denied" message
    expect(page.get_by_role("heading", name="Access Denied")).to_be_visible()

    # Expect to see the "Login" button
    login_button = page.get_by_role("button", name="Login")
    expect(login_button).to_be_visible()

    # 2. Click the login button and verify redirection
    login_button.click()

    # Expect the URL to be the login page
    expect(page).to_have_url("http://localhost:5173/login")

    # 3. Verify the login page content
    expect(page.get_by_role("heading", name="Login to ViralStock")).to_be_visible()

    # Expect to see the "Sign in with Google" button
    google_signin_button = page.get_by_role("button", name="Sign in with Google")
    expect(google_signin_button).to_be_visible()

    # 4. Take a screenshot for visual confirmation
    page.screenshot(path="jules-scratch/verification/auth_flow_verification.png")

def run():
    with sync_playwright() as p:
        browser = p.chromium.launch(headless=True)
        page = browser.new_page()
        test_logged_out_flow(page)
        browser.close()

run()
