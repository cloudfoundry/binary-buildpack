require 'spec_helper'

describe 'CF Binary Buildpack' do
  subject(:app) { Machete.deploy_app(app_name) }
  let(:browser) { Machete::Browser.new(app) }

  after do
    Machete::CF::DeleteApp.new.execute(app)
  end

  describe 'deploying a Ruby script' do
    subject(:app) { Machete.deploy_app(app_name, stack: 'lucid64') }
    let(:app_name) { 'webrick_app' }

    specify do
      browser.visit_path('/')

      expect(browser.body).to include('Hello, world!')
    end
  end
end
