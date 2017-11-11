FROM ruby:2.4.2-alpine

WORKDIR /usr/src/app

# COPY Gemfile ./Gemfile
# COPY Gemfile.lock ./Gemfile.lock
# RUN bundle install -j 20

COPY . .

CMD [ "ruby", "run.rb" ]
