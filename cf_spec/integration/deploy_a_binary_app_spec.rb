require 'spec_helper'

describe 'CF Binary Buildpack' do
  let(:browser) { Machete::Browser.new(app) }

  after do
    Machete::CF::DeleteApp.new.execute(app)
  end

  describe 'deploying a Ruby script' do
    let(:app_name) { 'webrick_app' }

    context 'when specifying a buildpack' do
      let(:app) { Machete.deploy_app(app_name, buildpack: 'binary-test-buildpack') }

      it 'deploys successfully' do
        expect(app).to be_running

        browser.visit_path('/')

        expect(browser.body).to include('Hello, world!')
      end
    end

    context 'without specifying a buildpack' do
      let(:app) { Machete.deploy_app(app_name) }

      it 'fails to stage' do
        expect(app).not_to be_running
        expect(app).to have_logged("An app was not successfully detected by any available buildpack")
      end
    end
  end
end
