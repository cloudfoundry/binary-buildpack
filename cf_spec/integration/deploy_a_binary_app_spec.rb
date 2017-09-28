require 'spec_helper'

describe 'CF Binary Buildpack' do
  let(:browser) { Machete::Browser.new(app) }

  after do
    Machete::CF::DeleteApp.new.execute(app)
  end

  describe 'deploying a Ruby script' do
    let(:app_name) { 'webrick_app' }

    context 'when specifying a buildpack' do
      let(:buildpack) { ENV.fetch('SHARED_HOST')=='true' ? 'binary_buildpack' : 'binary-test-buildpack' }
      let(:app) { Machete.deploy_app(app_name, buildpack: buildpack) }

      it 'deploys successfully' do
        expect(app).to be_running

        browser.visit_path('/')

        expect(browser.body).to include('Hello, world!')
      end
    end

    context 'without specifying a buildpack' do
      let(:app) { Machete.deploy_app(app_name, skip_verify_version: true) }

      it 'fails to stage' do
        expect(app).not_to be_running

        expect(app).to have_logged('None of the buildpacks detected a compatible application')
        end
      end
    end
  end
end
